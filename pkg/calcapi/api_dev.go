package calcapi

import (
	"context"

	"github.com/jinzhu/gorm"
)

func (svc *svc) CounterInc(ctx context.Context, in *CounterInc_Input) (*CounterInc_Output, error) {
	gap := in.Gap
	if gap == 0 {
		gap = 1
	}

	ret, err := svc.KeyValueFloatGet(ctx, &KeyValueFloatGet_Input{Key: in.Key})
	if err != nil {
		return nil, err
	}

	val := ret.Value + gap

	_, err = svc.KeyValueFloatSet(ctx, &KeyValueFloatSet_Input{Key: in.Key, Value: val})
	if err != nil {
		return nil, err
	}

	return &CounterInc_Output{Value: val}, nil
}

func (svc *svc) KeyValueStringGet(_ context.Context, in *KeyValueStringGet_Input) (*KeyValueStringGet_Output, error) {
	kv := DevKeyValueString{Key: in.Key}
	err := svc.db.Where(kv).First(&kv).Error
	if err != nil && gorm.IsRecordNotFoundError(err) {
		return &KeyValueStringGet_Output{Value: ""}, nil
	}
	if err != nil {
		return nil, err
	}

	return &KeyValueStringGet_Output{Value: kv.Value}, nil
}

func (svc *svc) KeyValueStringSet(ctx context.Context, in *KeyValueStringSet_Input) (*KeyValueStringSet_Output, error) {
	kv := DevKeyValueString{Key: in.Key, Value: in.Value}
	err := svc.db.Save(&kv).Error
	if err != nil {
		return nil, err
	}
	return &KeyValueStringSet_Output{}, nil
}

func (svc *svc) NumberSetIfBigger(ctx context.Context, in *NumberSetIfBigger_Input) (*NumberSetIfBigger_Output, error) {
	ret, err := svc.KeyValueFloatGet(ctx, &KeyValueFloatGet_Input{Key: in.Key})
	if err != nil {
		return nil, err
	}
	val := ret.Value
	if in.Value > val {
		val = in.Value
		_, err = svc.KeyValueFloatSet(ctx, &KeyValueFloatSet_Input{Key: in.Key, Value: val})
		if err != nil {
			return nil, err
		}
	}

	return &NumberSetIfBigger_Output{Value: val}, nil
}

func (svc *svc) KeyValueFloatGet(ctx context.Context, in *KeyValueFloatGet_Input) (*KeyValueFloatGet_Output, error) {
	kv := DevKeyValueFloat{Key: in.Key}
	err := svc.db.Where(kv).First(&kv).Error
	if err != nil && gorm.IsRecordNotFoundError(err) {
		return &KeyValueFloatGet_Output{Value: 0}, nil
	}
	if err != nil {
		return nil, err
	}

	return &KeyValueFloatGet_Output{Value: kv.Value}, nil
}

func (svc *svc) KeyValueFloatSet(ctx context.Context, in *KeyValueFloatSet_Input) (*KeyValueFloatSet_Output, error) {
	kv := DevKeyValueFloat{Key: in.Key, Value: in.Value}
	err := svc.db.Save(&kv).Error
	if err != nil {
		return nil, err
	}
	return &KeyValueFloatSet_Output{}, nil
}
