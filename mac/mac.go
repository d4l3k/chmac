package mac

import (
	"crypto/rand"
	"net"
	"os/exec"
)

func runCommand(name string, arg ...string) (string, error) {
	cmd := exec.Command(name, arg...)
	out, err := cmd.Output()
	if err != nil {
		return "", err
	}
	return string(out), nil
}

// SetMac sets the provided interface's mac address.
func SetMac(inter *net.Interface, addr net.HardwareAddr) error {
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

// SetRandMac sets the provided interface's mac address to a random one.
func SetRandMac(inter *net.Interface) error {
	mac, err := RandomMac()
	if err != nil {
		return err
	}
	if err := setMac(inter, mac); err != nil {
		return err
	}
	return nil
}
