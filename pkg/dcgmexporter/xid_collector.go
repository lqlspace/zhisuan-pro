/*
 * Copyright (c) 2024, VIRTAITECH CORPORATION.  All rights reserved.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package dcgmexporter

import (
	"fmt"
	"slices"

	"github.com/NVIDIA/go-dcgm/pkg/dcgm"
	"github.com/sirupsen/logrus"
)

type xidCollector struct {
	expCollector
}

func (c *xidCollector) GetMetrics() (MetricsByCounter, error) {
	return c.expCollector.getMetrics()
}

func NewXIDCollector(counters []Counter,
	hostname string,
	config *Config,
	fieldEntityGroupTypeSystemInfo FieldEntityGroupTypeSystemInfoItem) (Collector, error) {
	if !IsDCGMExpXIDErrorsCountEnabled(counters) {
		logrus.Error(dcgmExpXIDErrorsCount + " collector is disabled")
		return nil, fmt.Errorf(dcgmExpXIDErrorsCount + " collector is disabled")
	}

	collector := xidCollector{}
	collector.expCollector = newExpCollector(counters,
		hostname,
		[]dcgm.Short{dcgm.DCGM_FI_DEV_XID_ERRORS},
		config,
		fieldEntityGroupTypeSystemInfo)

	collector.counter = counters[slices.IndexFunc(counters, func(c Counter) bool {
		return c.FieldName == dcgmExpXIDErrorsCount
	})]

	collector.labelFiller = func(metricValueLabels map[string]string, entityValue int64) {
		metricValueLabels["xid"] = fmt.Sprint(entityValue)
	}

	collector.windowSize = config.XIDCountWindowSize

	return &collector, nil
}

func IsDCGMExpXIDErrorsCountEnabled(counters []Counter) bool {
	return slices.ContainsFunc(counters, func(c Counter) bool {
		return c.FieldName == dcgmExpXIDErrorsCount
	})
}
