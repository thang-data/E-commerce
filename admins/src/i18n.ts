import { createI18n, LocaleMessages } from "vue-i18n"
import locales from "@/locales"

function loadLocaleMessages(): LocaleMessages {
  return {
    ...locales,
  }
}

export default createI18n({
  locale: process.env.VUE_APP_I18N_LOCALE || "en",
  fallbackLocale: process.env.VUE_APP_I18N_FALLBACK_LOCALE || "en",
  messages: { ...locales },
})
