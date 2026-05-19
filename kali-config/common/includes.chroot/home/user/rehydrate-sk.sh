#!/usr/bin/env bash
set -euo pipefail

PIN="1234"
DEST_DIR=".ssh"

if ! command -v ssh-keygen >/dev/null 2>&1; then
  echo "Error: ssh-keygen not found" >&2
  exit 1
fi

if ! command -v expect >/dev/null 2>&1; then
  echo "Error: expect not found" >&2
  echo "Install it with: sudo apt install expect" >&2
  exit 1
fi

echo "Recovering resident keys"

EXPECT_STATUS=0
export PIN

expect <<'EOF' || EXPECT_STATUS=$?
set timeout -1
log_user 1

set pin $env(PIN)

spawn ssh-keygen -K -N ""

expect {
  -re "(?i)Enter PIN.*:" {
    send -- "$pin\r"
    exp_continue
  }
  -re "(?i)confirm passphrase.*:" {
    send -- "\r"
    exp_continue
  }
  -re "(?i)enter passphrase.*:" {
    send -- "\r"
    exp_continue
  }
  -re "(?i)touch your security key" {
    exp_continue
  }
  eof
}
catch wait result
set exit_status [lindex $result 3]
exit $exit_status
EOF

unset PIN

if [[ "$EXPECT_STATUS" -ne 0 ]]; then
  echo "Error: ssh-keygen -K failed" >&2
  exit "$EXPECT_STATUS"
fi

# Collect recovered key files for both supported types
shopt -s nullglob
ed25519_keys=(id_ed25519_sk_rk id_ed25519_sk_rk.pub)
ecdsa_keys=(id_ecdsa_sk_rk id_ecdsa_sk_rk.pub)
shopt -u nullglob

if [[ "${#ed25519_keys[@]}" -eq 0 && "${#ecdsa_keys[@]}" -eq 0 ]]; then
  echo "No keys were recovered." >&2
  exit 1
fi

# Install ed25519 keys (YubiKey / ed25519-capable devices)
if [[ "${#ed25519_keys[@]}" -gt 0 ]]; then
  mv ./id_ed25519_sk_rk     "$DEST_DIR/id_ed25519_sk"
  mv ./id_ed25519_sk_rk.pub "$DEST_DIR/id_ed25519_sk.pub"
  chmod 600 "$DEST_DIR/id_ed25519_sk"
  chmod 644 "$DEST_DIR/id_ed25519_sk.pub"
  echo "  ed25519-sk keys saved to: $DEST_DIR"
fi

# Install ecdsa keys (FIDO2 devices that only support ecdsa)
if [[ "${#ecdsa_keys[@]}" -gt 0 ]]; then
  mv ./id_ecdsa_sk_rk     "$DEST_DIR/id_ecdsa_sk"
  mv ./id_ecdsa_sk_rk.pub "$DEST_DIR/id_ecdsa_sk.pub"
  chmod 600 "$DEST_DIR/id_ecdsa_sk"
  chmod 644 "$DEST_DIR/id_ecdsa_sk.pub"
  echo "  ecdsa-sk keys saved to: $DEST_DIR"
fi

echo "Done."
