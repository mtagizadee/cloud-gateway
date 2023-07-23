package common

type PublicCertificate struct {
	AccessToken string `binding:"required"`
	ApplicationId int `binding:"required"`
  CertificateId string `binding:"required"`
  CreatedAt string `binding:"required"`
}

type BaseBodyDto[T comparable] struct {
	Certificate PublicCertificate
	Data T
}