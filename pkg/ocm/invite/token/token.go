// Copyright 2018-2023 CERN
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//
// In applying this license, CERN does not waive the privileges and immunities
// granted to it by virtue of its status as an Intergovernmental Organization
// or submit itself to any jurisdiction.

package token

import (
	"time"

	userpb "github.com/cs3org/go-cs3apis/cs3/identity/user/v1beta1"
	invitepb "github.com/cs3org/go-cs3apis/cs3/ocm/invite/v1beta1"
	typesv1beta1 "github.com/cs3org/go-cs3apis/cs3/types/v1beta1"
	"github.com/google/uuid"
	"github.com/pkg/errors"
)

// DefaultExpirationTime is the expiration time to be used when unspecified in the config.
const DefaultExpirationTime = "24h"

// CreateToken creates a InviteToken object for the userID indicated by userID.
func CreateToken(expiration string, userID *userpb.UserId) (*invitepb.InviteToken, error) {
	// Parse time of expiration
	duration, err := time.ParseDuration(expiration)
	if err != nil {
		return nil, errors.Wrap(err, "error parsing time of expiration")
	}

	tokenID := uuid.New().String()
	now := time.Now()
	expirationTime := now.Add(duration)

	token := invitepb.InviteToken{
		Token:  tokenID,
		UserId: userID,
		Expiration: &typesv1beta1.Timestamp{
			Seconds: uint64(expirationTime.Unix()),
			Nanos:   0,
		},
	}

	return &token, nil
}
