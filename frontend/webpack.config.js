module.exports = {
    module: {
        rules: [
            {
            test: /\.scss$/,
            use: [
                'vue-style-loader',
                'css-loader',
                'sass-loader'
            ]
            }
        ]
    },
    alias: {
        vue: 'vue/dist/vue.js'
    }
}