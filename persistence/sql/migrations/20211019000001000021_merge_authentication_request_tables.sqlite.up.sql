-- Migration generated by the command below; DO NOT EDIT.
-- hydra:generate hydra migrations gen

CREATE INDEX hydra_oauth2_oidc_request_id_idx ON hydra_oauth2_oidc (request_id);

