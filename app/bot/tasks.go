package bot

import (
	"github.com/YourSuzumiya/ACMBot/app/errs"
	"github.com/YourSuzumiya/ACMBot/app/manager"
	"github.com/YourSuzumiya/ACMBot/app/model"
)

type Task func(ctx *Context) error

// getHandlerFromParams nil -> []string
func getHandlerFromParams(ctx *Context) error {
	params := ctx.Params()
	if len(params) == 0 {
		return errs.ErrNoParam
	}

	var handles []string

	for _, param := range params {
		handle, ok := param.(string)
		if !ok {
			return errs.ErrBadParam
		}
		handles = append(handles, handle)
	}

	ctx.StepValue = handles
	return nil
}

// getCodeforcesUserByHandle []string -> *manager.CodeforcesUser
func getCodeforcesUserByHandle(ctx *Context) error {
	handles, ok := ctx.StepValue.([]string)
	if !ok {
		return errs.NewInternalError("错误的参数类型")
	}

	if len(handles) == 0 {
		ctx.SendError(errs.NewInternalError("getCodeforcesUserByHandle-handle数为0"))
		return errs.ErrBadBranch
	}

	if len(handles) > 1 {
		ctx.Send(Message{"太多handle惹，我只查询`" + handles[0] + "`的哦"})
	}

	user, err := manager.GetUpdatedCodeforcesUser(handles[0])
	if err != nil {
		return err
	}

	ctx.StepValue = user
	return nil
}

// getRenderedCodeforcesUserProfile *manager.CodeforcesUser -> []byte
func getRenderedCodeforcesUserProfile(ctx *Context) error {
	user, ok := ctx.StepValue.(*manager.CodeforcesUser)
	if !ok {
		return errs.NewInternalError("错误的参数类型")
	}

	pic, err := user.ToRenderProfileV2().ToImage()
	if err != nil {
		return err
	}

	ctx.StepValue = pic
	return nil
}

// getRenderedCodeforcesRatingChanges *manager.CodeforcesUser -> []byte
func getRenderedCodeforcesRatingChanges(ctx *Context) error {
	user, ok := ctx.StepValue.(*manager.CodeforcesUser)
	if !ok {
		return errs.NewInternalError("错误的参数类型")
	}

	pic, err := user.ToRenderRatingChanges().ToImage()
	if err != nil {
		return err
	}

	ctx.StepValue = pic
	return nil
}

// getRaceFromProvider model.RaceProvider -> []model.Race
func getRaceFromProvider(ctx *Context) error {
	provider, ok := ctx.StepValue.(model.RaceProvider)
	if !ok {
		return errs.NewInternalError("错误的参数类型")
	}

	race, err := provider()
	if err != nil {
		return err
	}

	ctx.StepValue = race
	return nil
}

// sendPicture []byte -> nil
func sendPicture(ctx *Context) error {
	pic, ok := ctx.StepValue.([]byte)
	if !ok {
		return errs.NewInternalError("错误的参数类型")
	}

	result := Message{pic}
	ctx.Send(result)
	ctx.StepValue = nil
	return nil
}

// []model.Race -> nil
func sendRace(ctx *Context) error {
	race, ok := ctx.StepValue.([]model.Race)
	if !ok {
		return errs.NewInternalError("错误的参数类型")
	}

	var result Message
	for _, v := range race {
		result = append(result, v.String())
	}
	ctx.Send(result)
	ctx.StepValue = nil
	return nil
}

func sendText(ctx *Context) error {
	text, ok := ctx.StepValue.(string)
	if !ok {
		return errs.NewInternalError("错误的参数类型")
	}
	ctx.Send(Message{text})
	ctx.StepValue = nil
	return nil
}
