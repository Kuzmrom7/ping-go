package pkg

import (
	"context"
	"log"
	"math"

	"github.com/Kuzmrom7/ping-go/config"
	"github.com/chromedp/cdproto/emulation"
	"github.com/chromedp/cdproto/page"
	"github.com/chromedp/chromedp"
)

func Run(cfg *config.Configurations) {
	ctx, cancel := chromedp.NewContext(context.Background())
	defer cancel()

	var imageBuf []byte
	if err := chromedp.Run(ctx, screenshotTasks(cfg.Url, &imageBuf)); err != nil {
		log.Println(err)
	}

	/* Send to telegram */
	err := SendPhoto(imageBuf, cfg.Bot)
	if err != nil {
		log.Println(err)
	} else {
		log.Println("Photo successfully sent")
	}
}

func screenshotTasks(url string, imageBuf *[]byte) chromedp.Tasks {
	return chromedp.Tasks{
		chromedp.Navigate(url),
		chromedp.ActionFunc(func(ctx context.Context) (err error) {
			_, _, contentSize, err := page.GetLayoutMetrics().Do(ctx)
			if err != nil {
				return err
			}

			width, height := int64(math.Ceil(contentSize.Width)), int64(math.Ceil(contentSize.Height))

			err = emulation.SetDeviceMetricsOverride(width, height, 1, false).
				WithScreenOrientation(&emulation.ScreenOrientation{
					Type:  emulation.OrientationTypePortraitPrimary,
					Angle: 0,
				}).
				Do(ctx)
			if err != nil {
				return err
			}

			*imageBuf, err = page.CaptureScreenshot().WithClip(&page.Viewport{
				X:      contentSize.X,
				Y:      contentSize.Y,
				Width:  contentSize.Width,
				Height: contentSize.Height,
				Scale:  1,
			}).WithQuality(100).Do(ctx)

			return err
		}),
	}
}
