/*
 * Copyright (C) 2019-2020 KallyDev
 * This program under GNU General Public License version 3.0, you
 * can redistribute it or modify it under the terms of the, see
 * the link below for more details
 *
 * https://github.com/kallydev/yogurt/blob/master/LICENSE
 */

package random_test

import (
	"fmt"
	"github.com/kallydev/yogurt/common/random"
	"testing"
)

func TestString(t *testing.T) {
	fmt.Println(random.String(8))
}
