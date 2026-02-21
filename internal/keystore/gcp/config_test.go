// Copyright 2022 - Hanzo AI, Inc. All rights reserved.
// Use of this source code is governed by the AGPLv3
// license that can be found in the LICENSE file.

package gcp

import (
	"testing"
)

func TestCredentials_MarshalJSON(t *testing.T) {
	for i, test := range marshalCredentialsTests {
		c := &Credentials{
			projectID: test.ProjectID,
			ClientID:  test.ClientID,
			Client:    test.ClientEmail,
			KeyID:     test.KeyID,
			Key:       test.Key,
		}
		got, err := c.MarshalJSON()
		if err != nil {
			t.Fatalf("Test %d: failed to marshal credentials: %v", i, err)
		}
		if s := string(got); s != test.JSON {
			t.Fatalf("Test %d:\n\ngot:\n%s\n\nwant:\n%s", i, s, test.JSON)
		}
	}
}

// testPrivateKey is a freshly generated RSA key used only for unit tests.
// It is NOT a real credential and has never been associated with any GCP project.
const testPrivateKey = "-----BEGIN PRIVATE KEY-----\nMIIEvwIBADANBgkqhkiG9w0BAQEFAASCBKkwggSlAgEAAoIBAQDSCs+qmKHdG/oZ\noOnzyAKILShCEJV6+6ZLWk7YiOSaZQNxA6shE68+Fp7WvGuYgI89TW8n0uYtBKRp\nVHG9/MxVshoJQSk9kQbj8tqdW8QwYem24bOK4p6mhKAVFMWhlaNzol+jeRouKImf\n2gL/W9n2dEdWDe7pMvX034HAnThNWlprZYlBNOrVludZLHznhuzlrarmj4mPeNF0\nOt3D1Dagd2xdZCN2uik9AYwbQ3m/dzCak72s6reT34e94yzzc7GgjuRAD/duVHg+\nHRh8kdoCtBsy2+Sw8RThHpn05ms11gJ3HunuvIKvCkpIYCc3PoeQKolWRmx6S54F\n1mFRPnePAgMBAAECggEAIU5Fj6+UOxNsnRO/BUD55c4DTDaE/XNZ8yjmx78tQK4z\n1X/Xau9MInRDy5klFpGMJRRaQ95VUV4jE/Cn+JBVdmE/b6g9ed29Uff4YmtiuKlN\n/gncBnHY5l6xY6ZfkDfCK1vb/nf9hsQ2znorb4KJFWNqzGJ5N9E03xdx47EKdHlN\nr+93nAm8rZWHCMiuSR4usBJEfBVBLWdW9rtcfVczfBGr1V4y1rntPGOH3bmERvEh\nhpvHqjJ+MoQQHkQxfdhSXqFU7cHMUE90N2eky7LNma1XjUrfpiZVgkrIGMDHcpH7\nVuRM4vDeqIQLF/Pm2x0xg8b41lkuRqCAEw47y8dNAQKBgQDyjPl1tGHOC9io2ExY\nxSlD5yDCh37/fuVlyG2b4rD3ZUMAhVqv7IbUfP0Cvp47EaefaE1hiAkgn2svmWlL\nibpkPRwZtlKfj2yKTJFp8gwYfOjOcf9GZ6K11LGKkGWZOmyPCn5wVTUdJPI6kzvY\n75+FcqQ3mHZ+AyzCZJqQGWtwnwKBgQDdsGB1zdN9LJ7LpoYZ15CCm5hsQUNzutXs\nmv4/5DMwkTMLTQyZdDU0RjTpgN+HldTgbUQ8EhsgMUt8Jlczivwa9pJuuUQXYBgZ\n61dagUUVXbcq/RtdB4hrDf7WJ2NzgDR9QaxnSlwyqPfZaeJQqz28hijKz7JBorRF\nw/XiUGHjEQKBgQDL8iHoCO6cVHWvM4Cuu8nA4N049Lmil87H0MnpY6mtvCWkkHz6\npAWTMi2EW0etH2E0tn0qXzpdhe6YrYROEdb6Zi/0psu+gtifSVaUIQRuqskOnegq\nCQZVoixO+K+VMf0KIabDScJJaUGMYzTValdV3FmcrcQknQLLyc/1Doq6vQKBgQC1\npuHThd2HEznnO8Nkh3tVWjPSeMPRn8XNA8/UAKRlZoUB/JeXYl56QBD5SO1gx0hQ\nwZgI9PkfLuVgjHYluPwNCABDFDIUSVxWWL4SXhyCfnEpm46Bczu4JRas2kemi+X2\nwQqI9KiJpmS41Qdp/Hcw9GmWjgvNtomrJQBk+mLLUQKBgQCQSLpFkg8ffluYlkNW\n+ywF7a4pv6xaeIri7N0Wrkd2/g8jhNGPAZtFW+yrKN22Ay2mhMddzaSPmipyy8D5\nX+H7LDhdXEQy5xDPg74s2CrN/uJwyd4StdWM0e69LfrOBLWc9CWmd2C8g0R7SlZm\nCh41SjMbqNLcuKP81eG4aiRnSg==\n-----END PRIVATE KEY-----\n"

var marshalCredentialsTests = []struct {
	ProjectID   string
	ClientID    string
	ClientEmail string
	KeyID       string
	Key         string
	JSON        string
}{
	{
		ProjectID:   "test-project-000000",
		ClientID:    "100000000000000000000",
		ClientEmail: "kes-test@test-project-000000.iam.gserviceaccount.com",
		KeyID:       "0000000000000000000000000000000000000000",
		Key:         testPrivateKey,
		JSON:        `{"type":"service_account","project_id":"test-project-000000","private_key_id":"0000000000000000000000000000000000000000","private_key":"-----BEGIN PRIVATE KEY-----\nMIIEvwIBADANBgkqhkiG9w0BAQEFAASCBKkwggSlAgEAAoIBAQDSCs+qmKHdG/oZ\noOnzyAKILShCEJV6+6ZLWk7YiOSaZQNxA6shE68+Fp7WvGuYgI89TW8n0uYtBKRp\nVHG9/MxVshoJQSk9kQbj8tqdW8QwYem24bOK4p6mhKAVFMWhlaNzol+jeRouKImf\n2gL/W9n2dEdWDe7pMvX034HAnThNWlprZYlBNOrVludZLHznhuzlrarmj4mPeNF0\nOt3D1Dagd2xdZCN2uik9AYwbQ3m/dzCak72s6reT34e94yzzc7GgjuRAD/duVHg+\nHRh8kdoCtBsy2+Sw8RThHpn05ms11gJ3HunuvIKvCkpIYCc3PoeQKolWRmx6S54F\n1mFRPnePAgMBAAECggEAIU5Fj6+UOxNsnRO/BUD55c4DTDaE/XNZ8yjmx78tQK4z\n1X/Xau9MInRDy5klFpGMJRRaQ95VUV4jE/Cn+JBVdmE/b6g9ed29Uff4YmtiuKlN\n/gncBnHY5l6xY6ZfkDfCK1vb/nf9hsQ2znorb4KJFWNqzGJ5N9E03xdx47EKdHlN\nr+93nAm8rZWHCMiuSR4usBJEfBVBLWdW9rtcfVczfBGr1V4y1rntPGOH3bmERvEh\nhpvHqjJ+MoQQHkQxfdhSXqFU7cHMUE90N2eky7LNma1XjUrfpiZVgkrIGMDHcpH7\nVuRM4vDeqIQLF/Pm2x0xg8b41lkuRqCAEw47y8dNAQKBgQDyjPl1tGHOC9io2ExY\nxSlD5yDCh37/fuVlyG2b4rD3ZUMAhVqv7IbUfP0Cvp47EaefaE1hiAkgn2svmWlL\nibpkPRwZtlKfj2yKTJFp8gwYfOjOcf9GZ6K11LGKkGWZOmyPCn5wVTUdJPI6kzvY\n75+FcqQ3mHZ+AyzCZJqQGWtwnwKBgQDdsGB1zdN9LJ7LpoYZ15CCm5hsQUNzutXs\nmv4/5DMwkTMLTQyZdDU0RjTpgN+HldTgbUQ8EhsgMUt8Jlczivwa9pJuuUQXYBgZ\n61dagUUVXbcq/RtdB4hrDf7WJ2NzgDR9QaxnSlwyqPfZaeJQqz28hijKz7JBorRF\nw/XiUGHjEQKBgQDL8iHoCO6cVHWvM4Cuu8nA4N049Lmil87H0MnpY6mtvCWkkHz6\npAWTMi2EW0etH2E0tn0qXzpdhe6YrYROEdb6Zi/0psu+gtifSVaUIQRuqskOnegq\nCQZVoixO+K+VMf0KIabDScJJaUGMYzTValdV3FmcrcQknQLLyc/1Doq6vQKBgQC1\npuHThd2HEznnO8Nkh3tVWjPSeMPRn8XNA8/UAKRlZoUB/JeXYl56QBD5SO1gx0hQ\nwZgI9PkfLuVgjHYluPwNCABDFDIUSVxWWL4SXhyCfnEpm46Bczu4JRas2kemi+X2\nwQqI9KiJpmS41Qdp/Hcw9GmWjgvNtomrJQBk+mLLUQKBgQCQSLpFkg8ffluYlkNW\n+ywF7a4pv6xaeIri7N0Wrkd2/g8jhNGPAZtFW+yrKN22Ay2mhMddzaSPmipyy8D5\nX+H7LDhdXEQy5xDPg74s2CrN/uJwyd4StdWM0e69LfrOBLWc9CWmd2C8g0R7SlZm\nCh41SjMbqNLcuKP81eG4aiRnSg==\n-----END PRIVATE KEY-----\n","client_email":"kes-test@test-project-000000.iam.gserviceaccount.com","client_id":"100000000000000000000","auth_uri":"https://accounts.google.com/o/oauth2/auth","token_uri":"https://accounts.google.com/o/oauth2/token","auth_provider_x509_cert_url":"https://www.googleapis.com/oauth2/v1/certs","client_x509_cert_url":"https://www.googleapis.com/robot/v1/metadata/x509/kes-test%40test-project-000000.iam.gserviceaccount.com"}`,
	},
	{
		ProjectID:   "test-project-000000",
		ClientID:    "100000000000000000000",
		ClientEmail: "kes-test@test-project-000000.iam.gserviceaccount.com",
		KeyID:       "0000000000000000000000000000000000000000",
		Key:         testPrivateKey,
		JSON:        `{"type":"service_account","project_id":"test-project-000000","private_key_id":"0000000000000000000000000000000000000000","private_key":"-----BEGIN PRIVATE KEY-----\nMIIEvwIBADANBgkqhkiG9w0BAQEFAASCBKkwggSlAgEAAoIBAQDSCs+qmKHdG/oZ\noOnzyAKILShCEJV6+6ZLWk7YiOSaZQNxA6shE68+Fp7WvGuYgI89TW8n0uYtBKRp\nVHG9/MxVshoJQSk9kQbj8tqdW8QwYem24bOK4p6mhKAVFMWhlaNzol+jeRouKImf\n2gL/W9n2dEdWDe7pMvX034HAnThNWlprZYlBNOrVludZLHznhuzlrarmj4mPeNF0\nOt3D1Dagd2xdZCN2uik9AYwbQ3m/dzCak72s6reT34e94yzzc7GgjuRAD/duVHg+\nHRh8kdoCtBsy2+Sw8RThHpn05ms11gJ3HunuvIKvCkpIYCc3PoeQKolWRmx6S54F\n1mFRPnePAgMBAAECggEAIU5Fj6+UOxNsnRO/BUD55c4DTDaE/XNZ8yjmx78tQK4z\n1X/Xau9MInRDy5klFpGMJRRaQ95VUV4jE/Cn+JBVdmE/b6g9ed29Uff4YmtiuKlN\n/gncBnHY5l6xY6ZfkDfCK1vb/nf9hsQ2znorb4KJFWNqzGJ5N9E03xdx47EKdHlN\nr+93nAm8rZWHCMiuSR4usBJEfBVBLWdW9rtcfVczfBGr1V4y1rntPGOH3bmERvEh\nhpvHqjJ+MoQQHkQxfdhSXqFU7cHMUE90N2eky7LNma1XjUrfpiZVgkrIGMDHcpH7\nVuRM4vDeqIQLF/Pm2x0xg8b41lkuRqCAEw47y8dNAQKBgQDyjPl1tGHOC9io2ExY\nxSlD5yDCh37/fuVlyG2b4rD3ZUMAhVqv7IbUfP0Cvp47EaefaE1hiAkgn2svmWlL\nibpkPRwZtlKfj2yKTJFp8gwYfOjOcf9GZ6K11LGKkGWZOmyPCn5wVTUdJPI6kzvY\n75+FcqQ3mHZ+AyzCZJqQGWtwnwKBgQDdsGB1zdN9LJ7LpoYZ15CCm5hsQUNzutXs\nmv4/5DMwkTMLTQyZdDU0RjTpgN+HldTgbUQ8EhsgMUt8Jlczivwa9pJuuUQXYBgZ\n61dagUUVXbcq/RtdB4hrDf7WJ2NzgDR9QaxnSlwyqPfZaeJQqz28hijKz7JBorRF\nw/XiUGHjEQKBgQDL8iHoCO6cVHWvM4Cuu8nA4N049Lmil87H0MnpY6mtvCWkkHz6\npAWTMi2EW0etH2E0tn0qXzpdhe6YrYROEdb6Zi/0psu+gtifSVaUIQRuqskOnegq\nCQZVoixO+K+VMf0KIabDScJJaUGMYzTValdV3FmcrcQknQLLyc/1Doq6vQKBgQC1\npuHThd2HEznnO8Nkh3tVWjPSeMPRn8XNA8/UAKRlZoUB/JeXYl56QBD5SO1gx0hQ\nwZgI9PkfLuVgjHYluPwNCABDFDIUSVxWWL4SXhyCfnEpm46Bczu4JRas2kemi+X2\nwQqI9KiJpmS41Qdp/Hcw9GmWjgvNtomrJQBk+mLLUQKBgQCQSLpFkg8ffluYlkNW\n+ywF7a4pv6xaeIri7N0Wrkd2/g8jhNGPAZtFW+yrKN22Ay2mhMddzaSPmipyy8D5\nX+H7LDhdXEQy5xDPg74s2CrN/uJwyd4StdWM0e69LfrOBLWc9CWmd2C8g0R7SlZm\nCh41SjMbqNLcuKP81eG4aiRnSg==\n-----END PRIVATE KEY-----\n","client_email":"kes-test@test-project-000000.iam.gserviceaccount.com","client_id":"100000000000000000000","auth_uri":"https://accounts.google.com/o/oauth2/auth","token_uri":"https://accounts.google.com/o/oauth2/token","auth_provider_x509_cert_url":"https://www.googleapis.com/oauth2/v1/certs","client_x509_cert_url":"https://www.googleapis.com/robot/v1/metadata/x509/kes-test%40test-project-000000.iam.gserviceaccount.com"}`,
	},
}
