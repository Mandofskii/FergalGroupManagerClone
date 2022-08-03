package globals

import "gopkg.in/telebot.v3"

var (
	StartKeyboard *telebot.ReplyMarkup = &telebot.ReplyMarkup{
		InlineKeyboard: [][]telebot.InlineButton{
			{
				telebot.InlineButton{
					Text: "🚀 افزودن ربات به گروه",
					URL:  "http://t.me/StrikerRbot?startgroup=new",
				},
			}, {
				telebot.InlineButton{
					Text: "📣 کانال آپدیت ها ",
					URL:  "https://t.me/fergaltm",
				},
				telebot.InlineButton{
					Text: "🌊 گروه پشتیبانی",
					URL:  "https://t.me/joinchat/LW3WU0wPSti1sUrGsP122g",
				},
			}, {
				telebot.InlineButton{
					Text: "🌟 لینکدونی",
					URL:  "https://t.me/LinkdoniFergal",
				},
				telebot.InlineButton{
					Text: "🔖 درباره تیم",
					Data: "about",
				},
			},
		},
	}
	AboutKeyboard *telebot.ReplyMarkup = &telebot.ReplyMarkup{
		InlineKeyboard: [][]telebot.InlineButton{
			{
				telebot.InlineButton{
					Text: "⚜️┈┅━ افزودن به گروه ━┅┈⚜️",
					URL:  "http://t.me/StrikerRbot?startgroup=new",
				},
			}, {
				telebot.InlineButton{
					Text: "کانال پشتیبانی ◽️ ",
					URL:  "http://t.me/StrikerRbot?startgroup=new",
				},
				telebot.InlineButton{
					Text: "◼️ گروه پشتیبانی",
					URL:  "https://t.me/joinchat/LW3WU0wPSti1sUrGsP122g",
				},
			}, {
				telebot.InlineButton{
					Text: "لینکدونی ◽️",
					URL:  "https://t.me/LinkdoniFergal",
				},
				telebot.InlineButton{
					Text: "◼️ خرید ممبر",
					URL:  "https://t.me/FergalShopBot",
				},
			}, {
				telebot.InlineButton{
					Text: "🎗┈┅━ درباره ربات ━┅┈🎗",
					Data: "about_bot",
				},
			},
		},
	}
	HelpKeyboard *telebot.ReplyMarkup = &telebot.ReplyMarkup{
		InlineKeyboard: [][]telebot.InlineButton{
			{
				telebot.InlineButton{
					Text: "╼ اد و جوین اجباری ╾ ",
					Data: "force_join_help",
				},
			}, {
				telebot.InlineButton{
					Text: "╼ فیلتر کلمات ",
					Data: "filter_help",
				},
				telebot.InlineButton{
					Text: "سکوت و حذف سکوت ╾",
					Data: "mute_help",
				},
			}, {
				telebot.InlineButton{
					Text: "╼ کاربر ویژه",
					Data: "vip_help",
				},
				telebot.InlineButton{
					Text: "بن و حذف بن ╾",
					Data: "ban_help",
				},
			}, {
				telebot.InlineButton{
					Text: "╼ ضد پورن",
					Data: "vip_help",
				},
				telebot.InlineButton{
					Text: "خوشامدگویی ╾",
					Data: "ban_help",
				},
			}, {
				telebot.InlineButton{
					Text: "╼ فان",
					Data: "vip_help",
				},
				telebot.InlineButton{
					Text: "قفل ها ╾",
					Data: "ban_help",
				},
			}, {
				telebot.InlineButton{
					Text: "╼ پاکسازی",
					Data: "vip_help",
				},
				telebot.InlineButton{
					Text: "ضد تبچی ╾",
					Data: "ban_help",
				},
			}, {
				telebot.InlineButton{
					Text: "╼ بی ادب و با ادب",
					Data: "vip_help",
				},
				telebot.InlineButton{
					Text: "ضد خیانت ╾",
					Data: "ban_help",
				},
			}, {
				telebot.InlineButton{
					Text: "╼ متفرقه",
					Data: "vip_help",
				},
				telebot.InlineButton{
					Text: "کانال و یوزرنیم مجاز ╾",
					Data: "ban_help",
				},
			}, {
				telebot.InlineButton{
					Text: "╼ پیام رگباری و ضد اسپم",
					Data: "vip_help",
				},
				telebot.InlineButton{
					Text: "آمار ╾",
					Data: "ban_help",
				},
			}, {
				telebot.InlineButton{
					Text: "╼ دسترسی مدیران",
					Data: "vip_help",
				},
				telebot.InlineButton{
					Text: "ارتقای مقام ╾",
					Data: "ban_help",
				},
			}, {
				telebot.InlineButton{
					Text: "╼ بازگشت به منوی اصلی ╾ ",
					Data: "vip_help",
				},
			},
		},
	}
)
