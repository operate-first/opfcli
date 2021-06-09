package api

import (
	"context"
	"fmt"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func (api *API) GetSecretDataString(nsname, secretname, fieldname string) (string, error) {
	client, err := api.Kubeclient()
	if err != nil {
		return "", err
	}

	secrets, err := client.CoreV1().Secrets(nsname).List(
		context.TODO(), metav1.ListOptions{
			FieldSelector: fmt.Sprintf(
				"metadata.name=%s",
				secretname,
			),
		})
	if err != nil {
		panic(err)
	}

	if count := len(secrets.Items); count != 1 {
		return "", fmt.Errorf(
			"failed to get secret %s/%s: expected 1, found %d",
			nsname,
			secretname,
			count,
		)
	}

	secret := secrets.Items[0]
	return string(secret.Data["ca.crt"]), nil
}
