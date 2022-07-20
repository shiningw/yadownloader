const path = require('path');
const webpack = require('webpack');
const MiniCssExtractPlugin = require("mini-css-extract-plugin");
const CompressionPlugin = require("compression-webpack-plugin");
const { VueLoaderPlugin } = require('vue-loader')

module.exports = {
  entry: {
    downloader: './src/main.js',
  },
  devtool: "source-map",
  output: {
    path: path.resolve(__dirname, './dist/js'),
    filename: '[name].js',
  },
  module: {
    rules: [
      {
        test: /\.s?[ac]ss$/i,
        use: [
          //'vue-style-loader',
          MiniCssExtractPlugin.loader,
          'css-loader',
          {
            loader: 'sass-loader',
            options: {
              additionalData: `
                @import "./src/css/_variables.scss";
              `
            }
          }
        ],
      },
      {
        test: /\.vue$/,
        use: 'vue-loader'
      },
      {
        test: /\.tsx?$/,
        use: 'ts-loader',
        exclude: /node_modules/,
      },
    ],
  },
  plugins: [
    new MiniCssExtractPlugin({
      filename: '../css/[name].css',
    }),
    new CompressionPlugin({
      include: /\.(js|map)$/,
      //deleteOriginalAssets: true,
    }),
    new VueLoaderPlugin()


  ],
  resolve: {
    extensions: [
      '.tsx',
      '.ts',
      '.js',
      '.jsx',
      '.vue',
      '.json',
    ],
    alias: {
      /* vue: 'vue/dist/vue.esm-bundler.js'*/
      assets: path.resolve(__dirname, 'img'),
      bulmaSrc: path.resolve(__dirname, 'node_modules/bulma/sass')
    },
  },
}
