package fail

import "github.com/mephistolie/chefbook-backend-common/responses/fail"

var (
	GrpcNameLength = fail.CreateGrpcClient(fail.TypeInvalidBody, "maximum first & last name lengths is 64 symbols")

	GrpcEmojiLength = fail.CreateGrpcClient(fail.TypeInvalidBody, "maximum emoji length is 25 symbols")
)
