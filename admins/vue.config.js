/* eslint @typescript-eslint/no-var-requires: "off" */
const path = require("path");
module.exports = {
  chainWebpack: (config) => {
    //this path is specific to my project
    config.resolve.alias.set("icons", path.resolve("src/assets/icons"));
    config.resolve.alias.set("images", path.resolve("src/assets/images"));
    // const svgRule = config.module.rule("svg");
    // svgRule.uses.clear();

    // svgRule
    //   .use("babel-loader")
    //   .loader("babel-loader")
    //   .end()
    //   .use("vue-svg-loader")
    //   .loader("vue-svg-loader");
  },

  // css: {
  //   loaderOptions: {
  //     sass: {
  //       prependData: `
  //         @import "@/scss/base.scss";
  //       `,
  //     },
  //   },
  // },

  pluginOptions: {
    i18n: {
      locale: "en",
      fallbackLocale: "en",
      localeDir: "locales",
      enableInSFC: false,
    },
  },
};
