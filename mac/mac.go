package mac

import (
	"crypto/rand"
	"net"
	"os"
	"os/exec"
)

func runCommand(name string, arg ...string) error {
	cmd := exec.Command(name, arg...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		return err
	}
	return nil
}

// SetMac sets the provided interface's mac address.
func SetMac(inter *net.Interface, addr net.HardwareAddr) error {
	if err := runCommand("ip", "link", "set", "dev", inter.Name, "down"); err != nil {
		return err
	}
	if err := runCommand("ip", "link", "set", "dev", inter.Name, "address", addr.String()); err != nil {
		return err
	}
	if err := runCommand("ip", "link", "set", "dev", inter.Name, "up"); err != nil {
		return err
	}
	return nil
}

// RandomMac returns a random mac address.
func RandomMac() (net.HardwareAddr, error) {
	addr := make([]byte, 6)
	if _, err := rand.Read(addr); err != nil {
		return nil, err
	}
	return net.HardwareAddr(addr), nil
}

// SetRandomMac sets the provided interface's mac address to a random one.
func SetRandomMac(inter *net.Interface) error {
	mac, err := RandomMac()
	if err != nil {
		return err
	}
	if err := SetMac(inter, mac); err != nil {
		return err
	}
	return nil
}
