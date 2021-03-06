package cluster

import (
	"strings"

	"github.com/giantswarm/microerror"

	"github.com/giantswarm/gsclientgen/v2/models"
	"github.com/giantswarm/gsctl/commands/errors"
)

func modifyClusterLabelsRequestFromArguments(labels []string) (*models.V5SetClusterLabelsRequest, error) {
	request := &models.V5SetClusterLabelsRequest{Labels: map[string]*string{}}

	for _, label := range labels {
		labelParts := strings.Split(strings.TrimSpace(label), "=")
		if len(labelParts) != 2 {
			return request, microerror.Maskf(errors.NoOpError, "malformed label change '%s' (single = required)", label)
		}
		if labelParts[0] == "" {
			return request, microerror.Maskf(errors.NoOpError, "malformed label change '%s' (empty key)", label)
		}
		if labelParts[1] == "" {
			request.Labels[labelParts[0]] = nil
		} else {
			request.Labels[labelParts[0]] = &labelParts[1]
		}
	}

	return request, nil
}
