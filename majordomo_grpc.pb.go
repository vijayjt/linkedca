// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package linkedca

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// MajordomoClient is the client API for Majordomo service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type MajordomoClient interface {
	// Login creates signs a given CSR and returns the certificate that will be
	// used for authentication.
	Login(ctx context.Context, in *LoginRequest, opts ...grpc.CallOption) (*LoginResponse, error)
	// GetRootCertificate returns the root certificate for a given fingerprint.
	GetRootCertificate(ctx context.Context, in *GetRootCertificateRequest, opts ...grpc.CallOption) (*GetRootCertificateResponse, error)
	// GetConfiguration returns the full configuration of an authority.
	GetConfiguration(ctx context.Context, in *ConfigurationRequest, opts ...grpc.CallOption) (*ConfigurationResponse, error)
	// CreateProvisioner adds a new provisioner to the majordomo authority and
	// returns the proto representation.
	CreateProvisioner(ctx context.Context, in *CreateProvisionerRequest, opts ...grpc.CallOption) (*Provisioner, error)
	// UpdateProvisioners updates a previously created provisioner.
	UpdateProvisioner(ctx context.Context, in *UpdateProvisionerRequest, opts ...grpc.CallOption) (*Provisioner, error)
	// DeleteProvisioner deletes a previously created provisioner.
	DeleteProvisioner(ctx context.Context, in *DeleteProvisionerRequest, opts ...grpc.CallOption) (*Provisioner, error)
	// CreateAdmin adds a new admin user to the majordomo authority. Admin users
	// can add or delete provisioners.
	CreateAdmin(ctx context.Context, in *CreateAdminRequest, opts ...grpc.CallOption) (*Admin, error)
	// UpdateAdmin updates a previously created admin.
	UpdateAdmin(ctx context.Context, in *UpdateAdminRequest, opts ...grpc.CallOption) (*Admin, error)
	// DeleteAdmin deletes a previously created admin user
	DeleteAdmin(ctx context.Context, in *DeleteAdminRequest, opts ...grpc.CallOption) (*Admin, error)
	// PostCertificate sends a signed X.509 certificate to majordomo.
	PostCertificate(ctx context.Context, in *CertificateRequest, opts ...grpc.CallOption) (*CertificateResponse, error)
	// PostSSHCertificate sends a signed SSH certificate to majordomo.
	PostSSHCertificate(ctx context.Context, in *SSHCertificateRequest, opts ...grpc.CallOption) (*SSHCertificateResponse, error)
	// RevokeCertificate marks an X.509 certificate as revoked.
	RevokeCertificate(ctx context.Context, in *RevokeCertificateRequest, opts ...grpc.CallOption) (*RevokeCertificateResponse, error)
	// RevokeSSHCertificate marks an SSH certificate as revoked.
	RevokeSSHCertificate(ctx context.Context, in *RevokeSSHCertificateRequest, opts ...grpc.CallOption) (*RevokeSSHCertificateResponse, error)
	// GetCertificateStatus returns the status of an X.509 certificate by serial.
	GetCertificateStatus(ctx context.Context, in *GetCertificateStatusRequest, opts ...grpc.CallOption) (*GetCertificateStatusResponse, error)
	// GetSSHCertificateStatus returns the status of an SSH certificate by serial.
	GetSSHCertificateStatus(ctx context.Context, in *GetSSHCertificateStatusRequest, opts ...grpc.CallOption) (*GetSSHCertificateStatusResponse, error)
}

type majordomoClient struct {
	cc grpc.ClientConnInterface
}

func NewMajordomoClient(cc grpc.ClientConnInterface) MajordomoClient {
	return &majordomoClient{cc}
}

func (c *majordomoClient) Login(ctx context.Context, in *LoginRequest, opts ...grpc.CallOption) (*LoginResponse, error) {
	out := new(LoginResponse)
	err := c.cc.Invoke(ctx, "/linkedca.Majordomo/Login", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *majordomoClient) GetRootCertificate(ctx context.Context, in *GetRootCertificateRequest, opts ...grpc.CallOption) (*GetRootCertificateResponse, error) {
	out := new(GetRootCertificateResponse)
	err := c.cc.Invoke(ctx, "/linkedca.Majordomo/GetRootCertificate", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *majordomoClient) GetConfiguration(ctx context.Context, in *ConfigurationRequest, opts ...grpc.CallOption) (*ConfigurationResponse, error) {
	out := new(ConfigurationResponse)
	err := c.cc.Invoke(ctx, "/linkedca.Majordomo/GetConfiguration", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *majordomoClient) CreateProvisioner(ctx context.Context, in *CreateProvisionerRequest, opts ...grpc.CallOption) (*Provisioner, error) {
	out := new(Provisioner)
	err := c.cc.Invoke(ctx, "/linkedca.Majordomo/CreateProvisioner", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *majordomoClient) UpdateProvisioner(ctx context.Context, in *UpdateProvisionerRequest, opts ...grpc.CallOption) (*Provisioner, error) {
	out := new(Provisioner)
	err := c.cc.Invoke(ctx, "/linkedca.Majordomo/UpdateProvisioner", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *majordomoClient) DeleteProvisioner(ctx context.Context, in *DeleteProvisionerRequest, opts ...grpc.CallOption) (*Provisioner, error) {
	out := new(Provisioner)
	err := c.cc.Invoke(ctx, "/linkedca.Majordomo/DeleteProvisioner", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *majordomoClient) CreateAdmin(ctx context.Context, in *CreateAdminRequest, opts ...grpc.CallOption) (*Admin, error) {
	out := new(Admin)
	err := c.cc.Invoke(ctx, "/linkedca.Majordomo/CreateAdmin", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *majordomoClient) UpdateAdmin(ctx context.Context, in *UpdateAdminRequest, opts ...grpc.CallOption) (*Admin, error) {
	out := new(Admin)
	err := c.cc.Invoke(ctx, "/linkedca.Majordomo/UpdateAdmin", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *majordomoClient) DeleteAdmin(ctx context.Context, in *DeleteAdminRequest, opts ...grpc.CallOption) (*Admin, error) {
	out := new(Admin)
	err := c.cc.Invoke(ctx, "/linkedca.Majordomo/DeleteAdmin", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *majordomoClient) PostCertificate(ctx context.Context, in *CertificateRequest, opts ...grpc.CallOption) (*CertificateResponse, error) {
	out := new(CertificateResponse)
	err := c.cc.Invoke(ctx, "/linkedca.Majordomo/PostCertificate", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *majordomoClient) PostSSHCertificate(ctx context.Context, in *SSHCertificateRequest, opts ...grpc.CallOption) (*SSHCertificateResponse, error) {
	out := new(SSHCertificateResponse)
	err := c.cc.Invoke(ctx, "/linkedca.Majordomo/PostSSHCertificate", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *majordomoClient) RevokeCertificate(ctx context.Context, in *RevokeCertificateRequest, opts ...grpc.CallOption) (*RevokeCertificateResponse, error) {
	out := new(RevokeCertificateResponse)
	err := c.cc.Invoke(ctx, "/linkedca.Majordomo/RevokeCertificate", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *majordomoClient) RevokeSSHCertificate(ctx context.Context, in *RevokeSSHCertificateRequest, opts ...grpc.CallOption) (*RevokeSSHCertificateResponse, error) {
	out := new(RevokeSSHCertificateResponse)
	err := c.cc.Invoke(ctx, "/linkedca.Majordomo/RevokeSSHCertificate", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *majordomoClient) GetCertificateStatus(ctx context.Context, in *GetCertificateStatusRequest, opts ...grpc.CallOption) (*GetCertificateStatusResponse, error) {
	out := new(GetCertificateStatusResponse)
	err := c.cc.Invoke(ctx, "/linkedca.Majordomo/GetCertificateStatus", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *majordomoClient) GetSSHCertificateStatus(ctx context.Context, in *GetSSHCertificateStatusRequest, opts ...grpc.CallOption) (*GetSSHCertificateStatusResponse, error) {
	out := new(GetSSHCertificateStatusResponse)
	err := c.cc.Invoke(ctx, "/linkedca.Majordomo/GetSSHCertificateStatus", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MajordomoServer is the server API for Majordomo service.
// All implementations must embed UnimplementedMajordomoServer
// for forward compatibility
type MajordomoServer interface {
	// Login creates signs a given CSR and returns the certificate that will be
	// used for authentication.
	Login(context.Context, *LoginRequest) (*LoginResponse, error)
	// GetRootCertificate returns the root certificate for a given fingerprint.
	GetRootCertificate(context.Context, *GetRootCertificateRequest) (*GetRootCertificateResponse, error)
	// GetConfiguration returns the full configuration of an authority.
	GetConfiguration(context.Context, *ConfigurationRequest) (*ConfigurationResponse, error)
	// CreateProvisioner adds a new provisioner to the majordomo authority and
	// returns the proto representation.
	CreateProvisioner(context.Context, *CreateProvisionerRequest) (*Provisioner, error)
	// UpdateProvisioners updates a previously created provisioner.
	UpdateProvisioner(context.Context, *UpdateProvisionerRequest) (*Provisioner, error)
	// DeleteProvisioner deletes a previously created provisioner.
	DeleteProvisioner(context.Context, *DeleteProvisionerRequest) (*Provisioner, error)
	// CreateAdmin adds a new admin user to the majordomo authority. Admin users
	// can add or delete provisioners.
	CreateAdmin(context.Context, *CreateAdminRequest) (*Admin, error)
	// UpdateAdmin updates a previously created admin.
	UpdateAdmin(context.Context, *UpdateAdminRequest) (*Admin, error)
	// DeleteAdmin deletes a previously created admin user
	DeleteAdmin(context.Context, *DeleteAdminRequest) (*Admin, error)
	// PostCertificate sends a signed X.509 certificate to majordomo.
	PostCertificate(context.Context, *CertificateRequest) (*CertificateResponse, error)
	// PostSSHCertificate sends a signed SSH certificate to majordomo.
	PostSSHCertificate(context.Context, *SSHCertificateRequest) (*SSHCertificateResponse, error)
	// RevokeCertificate marks an X.509 certificate as revoked.
	RevokeCertificate(context.Context, *RevokeCertificateRequest) (*RevokeCertificateResponse, error)
	// RevokeSSHCertificate marks an SSH certificate as revoked.
	RevokeSSHCertificate(context.Context, *RevokeSSHCertificateRequest) (*RevokeSSHCertificateResponse, error)
	// GetCertificateStatus returns the status of an X.509 certificate by serial.
	GetCertificateStatus(context.Context, *GetCertificateStatusRequest) (*GetCertificateStatusResponse, error)
	// GetSSHCertificateStatus returns the status of an SSH certificate by serial.
	GetSSHCertificateStatus(context.Context, *GetSSHCertificateStatusRequest) (*GetSSHCertificateStatusResponse, error)
	mustEmbedUnimplementedMajordomoServer()
}

// UnimplementedMajordomoServer must be embedded to have forward compatible implementations.
type UnimplementedMajordomoServer struct {
}

func (UnimplementedMajordomoServer) Login(context.Context, *LoginRequest) (*LoginResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Login not implemented")
}
func (UnimplementedMajordomoServer) GetRootCertificate(context.Context, *GetRootCertificateRequest) (*GetRootCertificateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetRootCertificate not implemented")
}
func (UnimplementedMajordomoServer) GetConfiguration(context.Context, *ConfigurationRequest) (*ConfigurationResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetConfiguration not implemented")
}
func (UnimplementedMajordomoServer) CreateProvisioner(context.Context, *CreateProvisionerRequest) (*Provisioner, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateProvisioner not implemented")
}
func (UnimplementedMajordomoServer) UpdateProvisioner(context.Context, *UpdateProvisionerRequest) (*Provisioner, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateProvisioner not implemented")
}
func (UnimplementedMajordomoServer) DeleteProvisioner(context.Context, *DeleteProvisionerRequest) (*Provisioner, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteProvisioner not implemented")
}
func (UnimplementedMajordomoServer) CreateAdmin(context.Context, *CreateAdminRequest) (*Admin, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateAdmin not implemented")
}
func (UnimplementedMajordomoServer) UpdateAdmin(context.Context, *UpdateAdminRequest) (*Admin, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateAdmin not implemented")
}
func (UnimplementedMajordomoServer) DeleteAdmin(context.Context, *DeleteAdminRequest) (*Admin, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteAdmin not implemented")
}
func (UnimplementedMajordomoServer) PostCertificate(context.Context, *CertificateRequest) (*CertificateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PostCertificate not implemented")
}
func (UnimplementedMajordomoServer) PostSSHCertificate(context.Context, *SSHCertificateRequest) (*SSHCertificateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PostSSHCertificate not implemented")
}
func (UnimplementedMajordomoServer) RevokeCertificate(context.Context, *RevokeCertificateRequest) (*RevokeCertificateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RevokeCertificate not implemented")
}
func (UnimplementedMajordomoServer) RevokeSSHCertificate(context.Context, *RevokeSSHCertificateRequest) (*RevokeSSHCertificateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RevokeSSHCertificate not implemented")
}
func (UnimplementedMajordomoServer) GetCertificateStatus(context.Context, *GetCertificateStatusRequest) (*GetCertificateStatusResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetCertificateStatus not implemented")
}
func (UnimplementedMajordomoServer) GetSSHCertificateStatus(context.Context, *GetSSHCertificateStatusRequest) (*GetSSHCertificateStatusResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetSSHCertificateStatus not implemented")
}
func (UnimplementedMajordomoServer) mustEmbedUnimplementedMajordomoServer() {}

// UnsafeMajordomoServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to MajordomoServer will
// result in compilation errors.
type UnsafeMajordomoServer interface {
	mustEmbedUnimplementedMajordomoServer()
}

func RegisterMajordomoServer(s grpc.ServiceRegistrar, srv MajordomoServer) {
	s.RegisterService(&Majordomo_ServiceDesc, srv)
}

func _Majordomo_Login_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LoginRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MajordomoServer).Login(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/linkedca.Majordomo/Login",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MajordomoServer).Login(ctx, req.(*LoginRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Majordomo_GetRootCertificate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetRootCertificateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MajordomoServer).GetRootCertificate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/linkedca.Majordomo/GetRootCertificate",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MajordomoServer).GetRootCertificate(ctx, req.(*GetRootCertificateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Majordomo_GetConfiguration_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ConfigurationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MajordomoServer).GetConfiguration(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/linkedca.Majordomo/GetConfiguration",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MajordomoServer).GetConfiguration(ctx, req.(*ConfigurationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Majordomo_CreateProvisioner_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateProvisionerRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MajordomoServer).CreateProvisioner(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/linkedca.Majordomo/CreateProvisioner",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MajordomoServer).CreateProvisioner(ctx, req.(*CreateProvisionerRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Majordomo_UpdateProvisioner_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateProvisionerRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MajordomoServer).UpdateProvisioner(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/linkedca.Majordomo/UpdateProvisioner",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MajordomoServer).UpdateProvisioner(ctx, req.(*UpdateProvisionerRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Majordomo_DeleteProvisioner_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteProvisionerRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MajordomoServer).DeleteProvisioner(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/linkedca.Majordomo/DeleteProvisioner",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MajordomoServer).DeleteProvisioner(ctx, req.(*DeleteProvisionerRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Majordomo_CreateAdmin_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateAdminRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MajordomoServer).CreateAdmin(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/linkedca.Majordomo/CreateAdmin",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MajordomoServer).CreateAdmin(ctx, req.(*CreateAdminRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Majordomo_UpdateAdmin_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateAdminRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MajordomoServer).UpdateAdmin(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/linkedca.Majordomo/UpdateAdmin",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MajordomoServer).UpdateAdmin(ctx, req.(*UpdateAdminRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Majordomo_DeleteAdmin_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteAdminRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MajordomoServer).DeleteAdmin(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/linkedca.Majordomo/DeleteAdmin",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MajordomoServer).DeleteAdmin(ctx, req.(*DeleteAdminRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Majordomo_PostCertificate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CertificateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MajordomoServer).PostCertificate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/linkedca.Majordomo/PostCertificate",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MajordomoServer).PostCertificate(ctx, req.(*CertificateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Majordomo_PostSSHCertificate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SSHCertificateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MajordomoServer).PostSSHCertificate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/linkedca.Majordomo/PostSSHCertificate",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MajordomoServer).PostSSHCertificate(ctx, req.(*SSHCertificateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Majordomo_RevokeCertificate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RevokeCertificateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MajordomoServer).RevokeCertificate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/linkedca.Majordomo/RevokeCertificate",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MajordomoServer).RevokeCertificate(ctx, req.(*RevokeCertificateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Majordomo_RevokeSSHCertificate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RevokeSSHCertificateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MajordomoServer).RevokeSSHCertificate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/linkedca.Majordomo/RevokeSSHCertificate",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MajordomoServer).RevokeSSHCertificate(ctx, req.(*RevokeSSHCertificateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Majordomo_GetCertificateStatus_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetCertificateStatusRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MajordomoServer).GetCertificateStatus(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/linkedca.Majordomo/GetCertificateStatus",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MajordomoServer).GetCertificateStatus(ctx, req.(*GetCertificateStatusRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Majordomo_GetSSHCertificateStatus_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetSSHCertificateStatusRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MajordomoServer).GetSSHCertificateStatus(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/linkedca.Majordomo/GetSSHCertificateStatus",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MajordomoServer).GetSSHCertificateStatus(ctx, req.(*GetSSHCertificateStatusRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Majordomo_ServiceDesc is the grpc.ServiceDesc for Majordomo service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Majordomo_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "linkedca.Majordomo",
	HandlerType: (*MajordomoServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Login",
			Handler:    _Majordomo_Login_Handler,
		},
		{
			MethodName: "GetRootCertificate",
			Handler:    _Majordomo_GetRootCertificate_Handler,
		},
		{
			MethodName: "GetConfiguration",
			Handler:    _Majordomo_GetConfiguration_Handler,
		},
		{
			MethodName: "CreateProvisioner",
			Handler:    _Majordomo_CreateProvisioner_Handler,
		},
		{
			MethodName: "UpdateProvisioner",
			Handler:    _Majordomo_UpdateProvisioner_Handler,
		},
		{
			MethodName: "DeleteProvisioner",
			Handler:    _Majordomo_DeleteProvisioner_Handler,
		},
		{
			MethodName: "CreateAdmin",
			Handler:    _Majordomo_CreateAdmin_Handler,
		},
		{
			MethodName: "UpdateAdmin",
			Handler:    _Majordomo_UpdateAdmin_Handler,
		},
		{
			MethodName: "DeleteAdmin",
			Handler:    _Majordomo_DeleteAdmin_Handler,
		},
		{
			MethodName: "PostCertificate",
			Handler:    _Majordomo_PostCertificate_Handler,
		},
		{
			MethodName: "PostSSHCertificate",
			Handler:    _Majordomo_PostSSHCertificate_Handler,
		},
		{
			MethodName: "RevokeCertificate",
			Handler:    _Majordomo_RevokeCertificate_Handler,
		},
		{
			MethodName: "RevokeSSHCertificate",
			Handler:    _Majordomo_RevokeSSHCertificate_Handler,
		},
		{
			MethodName: "GetCertificateStatus",
			Handler:    _Majordomo_GetCertificateStatus_Handler,
		},
		{
			MethodName: "GetSSHCertificateStatus",
			Handler:    _Majordomo_GetSSHCertificateStatus_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "majordomo.proto",
}
