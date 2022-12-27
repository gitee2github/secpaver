#!/bin/sh
# Copyright (c) Huawei Technologies Co., Ltd. 2020-2021. All rights reserved.
# secPaver is licensed under the Mulan PSL v2.
# You can use this software according to the terms and conditions of the Mulan PSL v2.
# You may obtain a copy of Mulan PSL v2 at:
#     http://license.coscl.org.cn/MulanPSL2
# THIS SOFTWARE IS PROVIDED ON AN "AS IS" BASIS, WITHOUT WARRANTIES OF ANY KIND, EITHER EXPRESS OR
# IMPLIED, INCLUDING BUT NOT LIMITED TO NON-INFRINGEMENT, MERCHANTABILITY OR FIT FOR A PARTICULAR
# PURPOSE.
# See the Mulan PSL v2 for more details.

source ./lib/test_lib.sh

TEST_NAME="TestPavdLogMod"

cleanup()
{
  rm -f $TEMP_LOG
}

test01()
{
  test_run "$TEST_NAME/$FUNCNAME"

  mod=$(stat -c %a /var/log/secpaver/pavd.log)

  if [ "$mod" != "600" ]; then
    echo "wrong file mod for /var/log/secpaver/pavd.log"
    test_fail "$TEST_NAME/$FUNCNAME"
    return
  fi

  test_pass "$TEST_NAME/$FUNCNAME"
}

test01

case_result

cleanup

