nix build --file ./postgres-tutorial.nix
QEMU_NET_OPTS='hostfwd=tcp::3000-:3000' result/bin/run-nixos-vm
