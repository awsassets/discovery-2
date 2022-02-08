package discovery

import (
	"context"
	"net"
	"time"
)

type Resolver struct {
	*net.Resolver
}

type Discovery struct {
	Target   string
	Address  net.IP
	Port     uint16
	Priority uint16
	Weight   uint16
}

// NewResolver returns resolver with given dns resolver
func NewResolver(resolver string) Resolver {
	return Resolver{
		&net.Resolver{
			PreferGo: true,
			Dial: func(ctx context.Context, network, address string) (net.Conn, error) {
				d := net.Dialer{
					Timeout: 2 * time.Second,
				}

				return d.DialContext(ctx, "udp", resolver)
			},
		},
	}
}

// LookupSrv wrap net.LookupSRV returns sorted records by priority
// and randomized by weight within a priority
func (r *Resolver) LookupSrv(ctx context.Context, service, proto, name string) ([]*net.SRV, error) {
	_, srv, err := r.LookupSRV(ctx, service, proto, name)
	if err != nil {
		return nil, err
	}

	return srv, err
}

// Lookup returns slice of ip address returned from given resolver
func (r *Resolver) Lookup(ctx context.Context, host string) ([]string, error) {
	ips, err := r.LookupHost(ctx, host)
	if err != nil {
		return nil, err
	}

	return ips, err
}

// Discover returns slice of discoverd service and error
func (r *Resolver) Discover(ctx context.Context, service, proto, name string) ([]Discovery, error) {
	var d []Discovery

	srv, err := r.LookupSrv(ctx, service, proto, name)
	if err != nil {
		return nil, err
	}

	for _, s := range srv {
		addr, err := r.getSingleIP(ctx, proto, s.Target)
		if err != nil {
			return nil, err
		}

		d = append(d, Discovery{s.Target, addr, s.Port, s.Priority, s.Weight})
	}

	return d, nil
}

// getSingleIP returns single ip from given host
func (r *Resolver) getSingleIP(ctx context.Context, network, host string) (net.IP, error) {
	ip, err := r.LookupIP(ctx, network, host)
	if err != nil {
		return nil, err
	}

	return ip[0], nil
}
