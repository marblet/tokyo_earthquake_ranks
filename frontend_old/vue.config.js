module.exports = {
  "transpileDependencies": [
    "vuetify"
  ],
  devServer: {
    proxy: {
      '/api': {
        target: 'https://tokyo-earthquake-ranks.df.r.appspot.com/'
        // target: 'http://localhost:8080'
      }
    }
  }
}