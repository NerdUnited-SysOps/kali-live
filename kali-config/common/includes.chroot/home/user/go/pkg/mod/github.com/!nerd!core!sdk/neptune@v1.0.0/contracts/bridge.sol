// SPDX-License-Identifier: MIT
pragma solidity 0.8.10;

contract Bridge {

    // Don't modify the order, as the sender & approver will be compacted into a single slot.
    struct Submission {
        uint256 value;
        address sender;
        bool approved;
    }

    struct SignedMessage {
        bytes32 key;
        address sender;
        uint256 amount;
    }

    event Received(bytes32 indexed key, address indexed sender, uint256 bridgedAmount, uint256 feePaid);
    event Approved(bytes32 indexed key, address indexed sender, uint256 bridgedAmount, address approver);
    event Notarized(address indexed sender, uint256 bridgedAmount, bytes32 nonce, bytes32 messageHash, bytes approvedMessage, bytes notarizeMessage, address notary);

    address private approver;
    address private notary;
    address private feeReceiver;
    mapping(bytes32 => Submission) private pending;
    mapping(bytes32 => bytes) private approvedMessage;
    bytes32 private domainSeparator;
    uint256 public fee;

    constructor(address _approver, address _notary, address _feeReceiver, uint256 _fee, uint256 _chainId) {
        require(_approver != address(0)); // dev: invalid approver
        require(_notary != address(0));   // dev: invalid notary
        
        approver    = _approver;
        notary      = _notary;
        feeReceiver = _feeReceiver;
        fee         = _fee;

        domainSeparator = keccak256(
            abi.encode(
                keccak256("EIP712Domain(string name,string version,uint256 chainId)"),
                keccak256("Neptune Bridge"), //name
                keccak256("0.0.1"), //version
                _chainId
            )
        );
    }


    receive() external payable { // msg.data is empty
        process();
    }

    fallback() external payable { // msg.data is NOT empty
        process();
    }

    function process()
    private {
        require(tx.origin == msg.sender);           // dev: sender must be the signer
        require(msg.value > fee);                   // dev: insufficient funds

        uint256 receiveAmount = msg.value-fee;
        require(receiveAmount > 0);                 //dev: invalid amount

        bytes32 key = keccak256(abi.encode(block.number, block.timestamp, msg.sender, receiveAmount));
        require(pending[key].sender == address(0)); // dev: submission already exists
        pending[key] = Submission(receiveAmount, msg.sender, false);

        emit Received(key, msg.sender, receiveAmount, fee);
    }

    function approve(bytes32 submission, bytes32 verifyHash, bytes memory sig)
    external {
        require(pending[submission].sender != address(0));       // dev: invalid submission

        if (msg.sender == approver) {
            require(pending[submission].approved == false);      // dev: already approved
            require(checkEncoding(verifyHash, sig, submission)); // dev: invalid signed message

            pending[submission].approved = true;
            approvedMessage[submission] = sig;

            emit Approved(submission, pending[submission].sender, pending[submission].value, approver);
        } else if (msg.sender == notary) {
            Submission memory s = pending[submission];
            bytes memory message = approvedMessage[submission];

            require(s.approved);                                  // dev: requires prior approval
            require(checkEncoding(verifyHash,  sig, submission)); // dev: invalid signed message

            delete pending[submission];
            delete approvedMessage[submission];

            (bool sent, bytes memory data) = address(feeReceiver).call{value: fee}("");
            require(sent);                                        // dev: failed to send fee
            (sent, data) = address(0).call{value: s.value}("");
            require(sent);                                        // dev: failed to burn eth

            emit Notarized(s.sender, s.value, submission, verifyHash, message, sig, notary);
        } else {
            revert();                                             // dev: invalid signer
        }
    }

    function checkEncoding(bytes32 verifyHash, bytes memory sig, bytes32 submission) 
    internal
    view
    returns(bool verified) {
        Submission memory sub = pending[submission];


        bytes32 hashToVerify = keccak256(
            abi.encode(keccak256("SignedMessage(bytes32 key,address sender,uint256 amount)"),submission,sub.sender,sub.value)
        );

        bytes32 domainSeparatorHash = keccak256(abi.encodePacked("\x19\x01", domainSeparator, hashToVerify));
        require(verifyHash == domainSeparatorHash); //dev: values do not match

        if (sub.approved == false) {
            verified = approver == recoverSigner(verifyHash, sig);
        } else if (sub.approved) {
            verified = notary == recoverSigner(verifyHash, sig);
        }

        return verified;
    }

    function splitSignature(bytes memory sig)
    internal
    pure
    returns (uint8 v, bytes32 r, bytes32 s) {
        require(sig.length == 65);

        assembly {
            r := mload(add(sig, 32))          // first 32 bytes, after the length prefix.
            s := mload(add(sig, 64))          // second 32 bytes.
            v := byte(0, mload(add(sig, 96))) // final byte (first byte of the next 32 bytes).
        }

        return (v, r, s);
    }

    function recoverSigner(bytes32 message, bytes memory sig)
    internal
    pure
    returns (address) {
        uint8 v;
        bytes32 r;
        bytes32 s;

        (v, r, s) = splitSignature(sig);

        return tryRecover(message, v, r, s);
    }

    function tryRecover(bytes32 hash, uint8 v, bytes32 r, bytes32 s)
    internal
    pure
    returns (address) {
        if (uint256(s) > 0x7FFFFFFFFFFFFFFFFFFFFFFFFFFFFFFF5D576E7357A4501DDFE92F46681B20A0) {
            return address(0);
        } else if (v != 27 && v != 28) {
            return address(0);
        }

        // If the signature is valid (and not malleable), return the signer address
        address signer = ecrecover(hash, v, r, s);
        if (signer == address(0)) {
            return address(0);
        }

        return signer;
    }

}
