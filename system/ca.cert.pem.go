package system

var caCertificate = []byte(`
-----BEGIN CERTIFICATE-----
MIIGoTCCBImgAwIBAgIJAMSnsbYqw4h7MA0GCSqGSIb3DQEBCwUAMIGRMQswCQYD
VQQGEwJFUzEOMAwGA1UECBMFU3BhaW4xEjAQBgNVBAcTCUJhcmNlbG9uYTERMA8G
A1UEChMIYXJhbWlyZXoxETAPBgNVBAsTCGFyYW1pcmV6MRQwEgYDVQQDEwthcmFt
aXJlei5lczEiMCAGCSqGSIb3DQEJARYTYWxiZXJ0b0BhcmFtaXJlei5lczAeFw0x
NTA1MjQxNTQzMTVaFw0yNTA1MjExNTQzMTVaMIGRMQswCQYDVQQGEwJFUzEOMAwG
A1UECBMFU3BhaW4xEjAQBgNVBAcTCUJhcmNlbG9uYTERMA8GA1UEChMIYXJhbWly
ZXoxETAPBgNVBAsTCGFyYW1pcmV6MRQwEgYDVQQDEwthcmFtaXJlei5lczEiMCAG
CSqGSIb3DQEJARYTYWxiZXJ0b0BhcmFtaXJlei5lczCCAiIwDQYJKoZIhvcNAQEB
BQADggIPADCCAgoCggIBANHJ3UvqImJ3e7pvWBJJ8tBOgstG1DBV6PE/Laa9F3MF
H3Xecd6BsWscQ6aRyHeq62S1ZAv2Hxseh25Y6rez/jHEIupJeZH9Gol1LkRmkJPa
ZtYEuF9gOT24MLlMLZcTIQpEH1U6styL4ayvwQkHTYQeE3MiEUdZa3HslMFkwNx3
EKy70TWXdxCGW9xTiC+on1LR5vQV3a/HbrzSbkWDX5e2lsRo7G7s5naNM62qkAT/
5LOLhmGRqmo0eqWeQ8k83AW+wp3Jyn8ZK3zdfDhb/F0sKjYHiaNBD1W5XRYrF3TO
KwBEV1FpRy4WIMP/12Qhm5rW77njhyKHNLsWO1vHJzkFmxioOqsbEsoypHDM/2Hf
xNqHPlIkwvsYFXmhgvOwAY7+njjtj4v+rGmq0FYWWujuHuT5DnT37S5zErjI1v+a
o7OxzpbQLt6drUHwBZTKXJK7V6tWM4BVmkiEPfBf0il0S1TD7z81mzXoLqqHgZAw
Ram5sGXlNzJvHqKfzz2ttkSeRhzmnmYYhYtUxWyuJ4OiJd8Qn5Pb4M0KJpL7SDX/
MpoZFjL2F8JWHm7Yw1zIE4AnqKzXXOQa1oq6oI+wG6x7Jck3ZgN9ccQbkOkcd9pV
12XUvqR/eMnkHspKk5pD+eFh20r3zQhGlotKES+oafUtRWe8GKEEPkcE42qDbSeV
AgMBAAGjgfkwgfYwHQYDVR0OBBYEFBgcR+S0SBHhRt3hgbEF+cIVqLqDMIHGBgNV
HSMEgb4wgbuAFBgcR+S0SBHhRt3hgbEF+cIVqLqDoYGXpIGUMIGRMQswCQYDVQQG
EwJFUzEOMAwGA1UECBMFU3BhaW4xEjAQBgNVBAcTCUJhcmNlbG9uYTERMA8GA1UE
ChMIYXJhbWlyZXoxETAPBgNVBAsTCGFyYW1pcmV6MRQwEgYDVQQDEwthcmFtaXJl
ei5lczEiMCAGCSqGSIb3DQEJARYTYWxiZXJ0b0BhcmFtaXJlei5lc4IJAMSnsbYq
w4h7MAwGA1UdEwQFMAMBAf8wDQYJKoZIhvcNAQELBQADggIBACRaWplIspqJCnqo
nmMWjtWYqeXxexZm/NtG5is0LDobpm1WQAIXvsNZ42YZZT3Ghvja7blCemljEyHO
TSdDDx0ulJLnj5du1WWIz/O7TRVPXmNiafMt2OezfExQzsQzW3sLkvRiKjcx21zY
QPmo6DB2XP6KpHeO/9MPJVJB4JK2xoX0pSMbcuqcx+uhv7SYh0oRTCx4oo+9KX33
LvnNaJw7/2eF+Vzv3DbnxSorSz2qLFDEf8RkozOjHYljci8RxUJiII3Rtu4B0sJY
K2WraO0RUVWJzktGTi7ksV9cjvDaudMoNdPzLGJh/NlaUxXMeDZelQvANZQD3lg5
G16t+UBOgnae9ena2b3g6fGzbaq8FI6SZOwQx9EWfg44NSvpcXI2Q92PEhsgOJpy
XsZRL169uqZ0cSlSoiCTCd2t0ujUwCwp/TB+wj8EUirhUlFtFhDd28jWUfMQ5n+a
lvyF0iSoymRmJesG2SG5i/wdX3s19vLIh1N/iLwpMSaMyiaC4ugXxKiSmQu6qEUJ
HM+UeqIJpvWfrYtIo2VGiGGS1A1sGqTYze2ED09JrgRreAqXFVmKNhFLEOHjONSm
Uib7+gx1HyPcfumuqyDQ4HwfcE0ZQtj3EtMJl/xV7lW2d1UBHhHaa3VnxlQHoThA
IVHSh2Tf5HijerI9atufoWvC69wm
-----END CERTIFICATE-----`);