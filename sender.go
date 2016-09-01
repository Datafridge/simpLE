package simpLE

import (
    "errors"
    //"fmt"
)

type sender struct {
    dev *device
    adv *advertisement
}

func (s *sender) device(string error) {

}

func (s *sender) Start() error {
    if s.dev == nil && s.adv == nil{
        return errors.New("device and advertisement not set")
    }
    if s.dev != nil && s.adv == nil{
        return errors.New("advertisement not set")
    }
    if s.dev == nil && s.adv != nil{
        return errors.New("device not set")
    }

    return nil
}
