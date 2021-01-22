// This file is auto-generated by internal/cmd/genheaders/main.go. DO NOT EDIT

package jwe

import (
	"context"
	"sort"
	"strconv"
	"sync"

	"github.com/lestrrat-go/jwx/internal/base64"
	"github.com/lestrrat-go/jwx/internal/json"
	"github.com/lestrrat-go/jwx/internal/pool"
	"github.com/lestrrat-go/jwx/jwa"
	"github.com/lestrrat-go/jwx/jwk"
	"github.com/pkg/errors"
)

const (
	AgreementPartyUInfoKey    = "apu"
	AgreementPartyVInfoKey    = "apv"
	AlgorithmKey              = "alg"
	CompressionKey            = "zip"
	ContentEncryptionKey      = "enc"
	ContentTypeKey            = "cty"
	CriticalKey               = "crit"
	EphemeralPublicKeyKey     = "epk"
	JWKKey                    = "jwk"
	JWKSetURLKey              = "jku"
	KeyIDKey                  = "kid"
	TypeKey                   = "typ"
	X509CertChainKey          = "x5c"
	X509CertThumbprintKey     = "x5t"
	X509CertThumbprintS256Key = "x5t#S256"
	X509URLKey                = "x5u"
)

// Headers describe a standard Header set.
type Headers interface {
	AgreementPartyUInfo() []byte
	AgreementPartyVInfo() []byte
	Algorithm() jwa.KeyEncryptionAlgorithm
	Compression() jwa.CompressionAlgorithm
	ContentEncryption() jwa.ContentEncryptionAlgorithm
	ContentType() string
	Critical() []string
	EphemeralPublicKey() jwk.Key
	JWK() jwk.Key
	JWKSetURL() string
	KeyID() string
	Type() string
	X509CertChain() []string
	X509CertThumbprint() string
	X509CertThumbprintS256() string
	X509URL() string
	Iterate(ctx context.Context) Iterator
	Walk(ctx context.Context, v Visitor) error
	AsMap(ctx context.Context) (map[string]interface{}, error)
	Get(string) (interface{}, bool)
	Set(string, interface{}) error
	Remove(string) error
	Encode() ([]byte, error)
	Decode([]byte) error
	// PrivateParams returns the map containing the non-standard ('private') parameters
	// in the associated header. WARNING: DO NOT USE PrivateParams()
	// IF YOU HAVE CONCURRENT CODE ACCESSING THEM. Use AsMap() to
	// get a copy of the entire header instead
	PrivateParams() map[string]interface{}
	Clone(context.Context) (Headers, error)
	Copy(context.Context, Headers) error
	Merge(context.Context, Headers) (Headers, error)
}

type stdHeaders struct {
	agreementPartyUInfo    []byte                          //
	agreementPartyVInfo    []byte                          //
	algorithm              *jwa.KeyEncryptionAlgorithm     //
	compression            *jwa.CompressionAlgorithm       //
	contentEncryption      *jwa.ContentEncryptionAlgorithm //
	contentType            *string                         //
	critical               []string                        //
	ephemeralPublicKey     jwk.Key                         //
	jwk                    jwk.Key                         //
	jwkSetURL              *string                         //
	keyID                  *string                         //
	typ                    *string                         //
	x509CertChain          []string                        //
	x509CertThumbprint     *string                         //
	x509CertThumbprintS256 *string                         //
	x509URL                *string                         //
	privateParams          map[string]interface{}
	mu                     *sync.RWMutex
}

type standardHeadersMarshalProxy struct {
	XagreementPartyUInfo    json.RawMessage                 `json:"apu,omitempty"`
	XagreementPartyVInfo    json.RawMessage                 `json:"apv,omitempty"`
	Xalgorithm              *jwa.KeyEncryptionAlgorithm     `json:"alg,omitempty"`
	Xcompression            *jwa.CompressionAlgorithm       `json:"zip,omitempty"`
	XcontentEncryption      *jwa.ContentEncryptionAlgorithm `json:"enc,omitempty"`
	XcontentType            *string                         `json:"cty,omitempty"`
	Xcritical               []string                        `json:"crit,omitempty"`
	XephemeralPublicKey     json.RawMessage                 `json:"epk,omitempty"`
	Xjwk                    json.RawMessage                 `json:"jwk,omitempty"`
	XjwkSetURL              *string                         `json:"jku,omitempty"`
	XkeyID                  *string                         `json:"kid,omitempty"`
	Xtyp                    *string                         `json:"typ,omitempty"`
	Xx509CertChain          []string                        `json:"x5c,omitempty"`
	Xx509CertThumbprint     *string                         `json:"x5t,omitempty"`
	Xx509CertThumbprintS256 *string                         `json:"x5t#S256,omitempty"`
	Xx509URL                *string                         `json:"x5u,omitempty"`
}

func NewHeaders() Headers {
	return &stdHeaders{
		mu:            &sync.RWMutex{},
		privateParams: map[string]interface{}{},
	}
}

func (h *stdHeaders) AgreementPartyUInfo() []byte {
	h.mu.RLock()
	defer h.mu.RUnlock()
	return h.agreementPartyUInfo
}

func (h *stdHeaders) AgreementPartyVInfo() []byte {
	h.mu.RLock()
	defer h.mu.RUnlock()
	return h.agreementPartyVInfo
}

func (h *stdHeaders) Algorithm() jwa.KeyEncryptionAlgorithm {
	h.mu.RLock()
	defer h.mu.RUnlock()
	if h.algorithm == nil {
		return ""
	}
	return *(h.algorithm)
}

func (h *stdHeaders) Compression() jwa.CompressionAlgorithm {
	h.mu.RLock()
	defer h.mu.RUnlock()
	if h.compression == nil {
		return ""
	}
	return *(h.compression)
}

func (h *stdHeaders) ContentEncryption() jwa.ContentEncryptionAlgorithm {
	h.mu.RLock()
	defer h.mu.RUnlock()
	if h.contentEncryption == nil {
		return ""
	}
	return *(h.contentEncryption)
}

func (h *stdHeaders) ContentType() string {
	h.mu.RLock()
	defer h.mu.RUnlock()
	if h.contentType == nil {
		return ""
	}
	return *(h.contentType)
}

func (h *stdHeaders) Critical() []string {
	h.mu.RLock()
	defer h.mu.RUnlock()
	return h.critical
}

func (h *stdHeaders) EphemeralPublicKey() jwk.Key {
	h.mu.RLock()
	defer h.mu.RUnlock()
	return h.ephemeralPublicKey
}

func (h *stdHeaders) JWK() jwk.Key {
	h.mu.RLock()
	defer h.mu.RUnlock()
	return h.jwk
}

func (h *stdHeaders) JWKSetURL() string {
	h.mu.RLock()
	defer h.mu.RUnlock()
	if h.jwkSetURL == nil {
		return ""
	}
	return *(h.jwkSetURL)
}

func (h *stdHeaders) KeyID() string {
	h.mu.RLock()
	defer h.mu.RUnlock()
	if h.keyID == nil {
		return ""
	}
	return *(h.keyID)
}

func (h *stdHeaders) Type() string {
	h.mu.RLock()
	defer h.mu.RUnlock()
	if h.typ == nil {
		return ""
	}
	return *(h.typ)
}

func (h *stdHeaders) X509CertChain() []string {
	h.mu.RLock()
	defer h.mu.RUnlock()
	return h.x509CertChain
}

func (h *stdHeaders) X509CertThumbprint() string {
	h.mu.RLock()
	defer h.mu.RUnlock()
	if h.x509CertThumbprint == nil {
		return ""
	}
	return *(h.x509CertThumbprint)
}

func (h *stdHeaders) X509CertThumbprintS256() string {
	h.mu.RLock()
	defer h.mu.RUnlock()
	if h.x509CertThumbprintS256 == nil {
		return ""
	}
	return *(h.x509CertThumbprintS256)
}

func (h *stdHeaders) X509URL() string {
	h.mu.RLock()
	defer h.mu.RUnlock()
	if h.x509URL == nil {
		return ""
	}
	return *(h.x509URL)
}

func (h *stdHeaders) iterate(ctx context.Context, ch chan *HeaderPair) {
	defer close(ch)
	h.mu.RLock()
	defer h.mu.RUnlock()
	var pairs []*HeaderPair
	if h.agreementPartyUInfo != nil {
		pairs = append(pairs, &HeaderPair{Key: AgreementPartyUInfoKey, Value: h.agreementPartyUInfo})
	}
	if h.agreementPartyVInfo != nil {
		pairs = append(pairs, &HeaderPair{Key: AgreementPartyVInfoKey, Value: h.agreementPartyVInfo})
	}
	if h.algorithm != nil {
		pairs = append(pairs, &HeaderPair{Key: AlgorithmKey, Value: *(h.algorithm)})
	}
	if h.compression != nil {
		pairs = append(pairs, &HeaderPair{Key: CompressionKey, Value: *(h.compression)})
	}
	if h.contentEncryption != nil {
		pairs = append(pairs, &HeaderPair{Key: ContentEncryptionKey, Value: *(h.contentEncryption)})
	}
	if h.contentType != nil {
		pairs = append(pairs, &HeaderPair{Key: ContentTypeKey, Value: *(h.contentType)})
	}
	if h.critical != nil {
		pairs = append(pairs, &HeaderPair{Key: CriticalKey, Value: h.critical})
	}
	if h.ephemeralPublicKey != nil {
		pairs = append(pairs, &HeaderPair{Key: EphemeralPublicKeyKey, Value: h.ephemeralPublicKey})
	}
	if h.jwk != nil {
		pairs = append(pairs, &HeaderPair{Key: JWKKey, Value: h.jwk})
	}
	if h.jwkSetURL != nil {
		pairs = append(pairs, &HeaderPair{Key: JWKSetURLKey, Value: *(h.jwkSetURL)})
	}
	if h.keyID != nil {
		pairs = append(pairs, &HeaderPair{Key: KeyIDKey, Value: *(h.keyID)})
	}
	if h.typ != nil {
		pairs = append(pairs, &HeaderPair{Key: TypeKey, Value: *(h.typ)})
	}
	if h.x509CertChain != nil {
		pairs = append(pairs, &HeaderPair{Key: X509CertChainKey, Value: h.x509CertChain})
	}
	if h.x509CertThumbprint != nil {
		pairs = append(pairs, &HeaderPair{Key: X509CertThumbprintKey, Value: *(h.x509CertThumbprint)})
	}
	if h.x509CertThumbprintS256 != nil {
		pairs = append(pairs, &HeaderPair{Key: X509CertThumbprintS256Key, Value: *(h.x509CertThumbprintS256)})
	}
	if h.x509URL != nil {
		pairs = append(pairs, &HeaderPair{Key: X509URLKey, Value: *(h.x509URL)})
	}
	for k, v := range h.privateParams {
		pairs = append(pairs, &HeaderPair{Key: k, Value: v})
	}
	for _, pair := range pairs {
		select {
		case <-ctx.Done():
			return
		case ch <- pair:
		}
	}
}

func (h *stdHeaders) PrivateParams() map[string]interface{} {
	h.mu.RLock()
	defer h.mu.RUnlock()
	return h.privateParams
}

func (h *stdHeaders) Get(name string) (interface{}, bool) {
	h.mu.RLock()
	defer h.mu.RUnlock()
	switch name {
	case AgreementPartyUInfoKey:
		if h.agreementPartyUInfo == nil {
			return nil, false
		}
		return h.agreementPartyUInfo, true
	case AgreementPartyVInfoKey:
		if h.agreementPartyVInfo == nil {
			return nil, false
		}
		return h.agreementPartyVInfo, true
	case AlgorithmKey:
		if h.algorithm == nil {
			return nil, false
		}
		return *(h.algorithm), true
	case CompressionKey:
		if h.compression == nil {
			return nil, false
		}
		return *(h.compression), true
	case ContentEncryptionKey:
		if h.contentEncryption == nil {
			return nil, false
		}
		return *(h.contentEncryption), true
	case ContentTypeKey:
		if h.contentType == nil {
			return nil, false
		}
		return *(h.contentType), true
	case CriticalKey:
		if h.critical == nil {
			return nil, false
		}
		return h.critical, true
	case EphemeralPublicKeyKey:
		if h.ephemeralPublicKey == nil {
			return nil, false
		}
		return h.ephemeralPublicKey, true
	case JWKKey:
		if h.jwk == nil {
			return nil, false
		}
		return h.jwk, true
	case JWKSetURLKey:
		if h.jwkSetURL == nil {
			return nil, false
		}
		return *(h.jwkSetURL), true
	case KeyIDKey:
		if h.keyID == nil {
			return nil, false
		}
		return *(h.keyID), true
	case TypeKey:
		if h.typ == nil {
			return nil, false
		}
		return *(h.typ), true
	case X509CertChainKey:
		if h.x509CertChain == nil {
			return nil, false
		}
		return h.x509CertChain, true
	case X509CertThumbprintKey:
		if h.x509CertThumbprint == nil {
			return nil, false
		}
		return *(h.x509CertThumbprint), true
	case X509CertThumbprintS256Key:
		if h.x509CertThumbprintS256 == nil {
			return nil, false
		}
		return *(h.x509CertThumbprintS256), true
	case X509URLKey:
		if h.x509URL == nil {
			return nil, false
		}
		return *(h.x509URL), true
	default:
		v, ok := h.privateParams[name]
		return v, ok
	}
}

func (h *stdHeaders) Set(name string, value interface{}) error {
	h.mu.Lock()
	defer h.mu.Unlock()
	switch name {
	case AgreementPartyUInfoKey:
		if v, ok := value.([]byte); ok {
			h.agreementPartyUInfo = v
			return nil
		}
		return errors.Errorf(`invalid value for %s key: %T`, AgreementPartyUInfoKey, value)
	case AgreementPartyVInfoKey:
		if v, ok := value.([]byte); ok {
			h.agreementPartyVInfo = v
			return nil
		}
		return errors.Errorf(`invalid value for %s key: %T`, AgreementPartyVInfoKey, value)
	case AlgorithmKey:
		if v, ok := value.(jwa.KeyEncryptionAlgorithm); ok {
			h.algorithm = &v
			return nil
		}
		return errors.Errorf(`invalid value for %s key: %T`, AlgorithmKey, value)
	case CompressionKey:
		if v, ok := value.(jwa.CompressionAlgorithm); ok {
			h.compression = &v
			return nil
		}
		return errors.Errorf(`invalid value for %s key: %T`, CompressionKey, value)
	case ContentEncryptionKey:
		if v, ok := value.(jwa.ContentEncryptionAlgorithm); ok {
			if v == "" {
				return errors.New(`"enc" field cannot be an empty string`)
			}
			h.contentEncryption = &v
			return nil
		}
		return errors.Errorf(`invalid value for %s key: %T`, ContentEncryptionKey, value)
	case ContentTypeKey:
		if v, ok := value.(string); ok {
			h.contentType = &v
			return nil
		}
		return errors.Errorf(`invalid value for %s key: %T`, ContentTypeKey, value)
	case CriticalKey:
		if v, ok := value.([]string); ok {
			h.critical = v
			return nil
		}
		return errors.Errorf(`invalid value for %s key: %T`, CriticalKey, value)
	case EphemeralPublicKeyKey:
		if v, ok := value.(jwk.Key); ok {
			h.ephemeralPublicKey = v
			return nil
		}
		return errors.Errorf(`invalid value for %s key: %T`, EphemeralPublicKeyKey, value)
	case JWKKey:
		if v, ok := value.(jwk.Key); ok {
			h.jwk = v
			return nil
		}
		return errors.Errorf(`invalid value for %s key: %T`, JWKKey, value)
	case JWKSetURLKey:
		if v, ok := value.(string); ok {
			h.jwkSetURL = &v
			return nil
		}
		return errors.Errorf(`invalid value for %s key: %T`, JWKSetURLKey, value)
	case KeyIDKey:
		if v, ok := value.(string); ok {
			h.keyID = &v
			return nil
		}
		return errors.Errorf(`invalid value for %s key: %T`, KeyIDKey, value)
	case TypeKey:
		if v, ok := value.(string); ok {
			h.typ = &v
			return nil
		}
		return errors.Errorf(`invalid value for %s key: %T`, TypeKey, value)
	case X509CertChainKey:
		if v, ok := value.([]string); ok {
			h.x509CertChain = v
			return nil
		}
		return errors.Errorf(`invalid value for %s key: %T`, X509CertChainKey, value)
	case X509CertThumbprintKey:
		if v, ok := value.(string); ok {
			h.x509CertThumbprint = &v
			return nil
		}
		return errors.Errorf(`invalid value for %s key: %T`, X509CertThumbprintKey, value)
	case X509CertThumbprintS256Key:
		if v, ok := value.(string); ok {
			h.x509CertThumbprintS256 = &v
			return nil
		}
		return errors.Errorf(`invalid value for %s key: %T`, X509CertThumbprintS256Key, value)
	case X509URLKey:
		if v, ok := value.(string); ok {
			h.x509URL = &v
			return nil
		}
		return errors.Errorf(`invalid value for %s key: %T`, X509URLKey, value)
	default:
		if h.privateParams == nil {
			h.privateParams = map[string]interface{}{}
		}
		h.privateParams[name] = value
	}
	return nil
}

func (h *stdHeaders) Remove(key string) error {
	h.mu.Lock()
	defer h.mu.Unlock()
	switch key {
	case AgreementPartyUInfoKey:
		h.agreementPartyUInfo = nil
	case AgreementPartyVInfoKey:
		h.agreementPartyVInfo = nil
	case AlgorithmKey:
		h.algorithm = nil
	case CompressionKey:
		h.compression = nil
	case ContentEncryptionKey:
		h.contentEncryption = nil
	case ContentTypeKey:
		h.contentType = nil
	case CriticalKey:
		h.critical = nil
	case EphemeralPublicKeyKey:
		h.ephemeralPublicKey = nil
	case JWKKey:
		h.jwk = nil
	case JWKSetURLKey:
		h.jwkSetURL = nil
	case KeyIDKey:
		h.keyID = nil
	case TypeKey:
		h.typ = nil
	case X509CertChainKey:
		h.x509CertChain = nil
	case X509CertThumbprintKey:
		h.x509CertThumbprint = nil
	case X509CertThumbprintS256Key:
		h.x509CertThumbprintS256 = nil
	case X509URLKey:
		h.x509URL = nil
	default:
		delete(h.privateParams, key)
	}
	return nil
}

func (h *stdHeaders) UnmarshalJSON(buf []byte) error {
	h.mu.Lock()
	defer h.mu.Unlock()
	var proxy standardHeadersMarshalProxy
	if err := json.Unmarshal(buf, &proxy); err != nil {
		return errors.Wrap(err, `failed to unmarshal headers`)
	}

	h.jwk = nil
	if jwkField := proxy.Xjwk; len(jwkField) > 0 {
		jwkKey, err := jwk.ParseKey([]byte(proxy.Xjwk))
		if err != nil {
			return errors.Wrap(err, `failed to parse jwk field`)
		}
		h.jwk = jwkKey
	}

	h.ephemeralPublicKey = nil
	if epkField := proxy.XephemeralPublicKey; len(epkField) > 0 {
		epk, err := jwk.ParseKey([]byte(proxy.XephemeralPublicKey))
		if err != nil {
			return errors.Wrap(err, `failed to parse epk field`)
		}
		h.ephemeralPublicKey = epk
	}

	h.agreementPartyUInfo = nil
	if v := proxy.XagreementPartyUInfo; len(v) > 0 {
		decoded, err := base64.Decode(v)
		if err != nil {
			return errors.Wrap(err, `failed to decode base64`)
		}
		h.agreementPartyUInfo = decoded
	}

	h.agreementPartyVInfo = nil
	if v := proxy.XagreementPartyVInfo; len(v) > 0 {
		decoded, err := base64.Decode(v)
		if err != nil {
			return errors.Wrap(err, `failed to decode base64`)
		}
		h.agreementPartyVInfo = decoded
	}
	h.algorithm = proxy.Xalgorithm
	h.compression = proxy.Xcompression
	h.contentEncryption = proxy.XcontentEncryption
	h.contentType = proxy.XcontentType
	h.critical = proxy.Xcritical
	h.jwkSetURL = proxy.XjwkSetURL
	h.keyID = proxy.XkeyID
	h.typ = proxy.Xtyp
	h.x509CertChain = proxy.Xx509CertChain
	h.x509CertThumbprint = proxy.Xx509CertThumbprint
	h.x509CertThumbprintS256 = proxy.Xx509CertThumbprintS256
	h.x509URL = proxy.Xx509URL
	var m map[string]interface{}
	if err := json.Unmarshal(buf, &m); err != nil {
		return errors.Wrap(err, `failed to parse privsate parameters`)
	}
	delete(m, AgreementPartyUInfoKey)
	delete(m, AgreementPartyVInfoKey)
	delete(m, AlgorithmKey)
	delete(m, CompressionKey)
	delete(m, ContentEncryptionKey)
	delete(m, ContentTypeKey)
	delete(m, CriticalKey)
	delete(m, EphemeralPublicKeyKey)
	delete(m, JWKKey)
	delete(m, JWKSetURLKey)
	delete(m, KeyIDKey)
	delete(m, TypeKey)
	delete(m, X509CertChainKey)
	delete(m, X509CertThumbprintKey)
	delete(m, X509CertThumbprintS256Key)
	delete(m, X509URLKey)
	h.privateParams = m
	return nil
}

func (h stdHeaders) MarshalJSON() ([]byte, error) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	data := make(map[string]interface{})
	fields := make([]string, 0, 16)
	for iter := h.Iterate(ctx); iter.Next(ctx); {
		pair := iter.Pair()
		fields = append(fields, pair.Key.(string))
		data[pair.Key.(string)] = pair.Value
	}

	sort.Strings(fields)
	buf := pool.GetBytesBuffer()
	defer pool.ReleaseBytesBuffer(buf)
	buf.WriteByte('{')
	l := len(fields)
	enc := json.NewEncoder(buf)
	for i, f := range fields {
		buf.WriteString(strconv.Quote(f))
		buf.WriteByte(':')
		v := data[f]
		switch v := v.(type) {
		case []byte:
			enc.Encode(base64.EncodeToString(v))
		default:
			enc.Encode(v)
		}

		if i < l-1 {
			buf.WriteByte(',')
		}
	}
	buf.WriteByte('}')
	ret := make([]byte, buf.Len())
	copy(ret, buf.Bytes())
	return ret, nil
}
