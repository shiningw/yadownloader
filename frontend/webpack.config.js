const path = require('path');
const webpack = require('webpack');
const { VueLoaderPlugin } = require('vue-loader')
//const MiniCssExtractPlugin = require("mini-css-extract-plugin");
const CompressionPlugin = require("compression-webpack-plugin");


module.exports = {
  /* experiments: {
     asset: true
   },*/
  entry: {
    downloader: './src/main.js',
  },
  devtool: "source-map",
  output: {
    path: path.resolve(__dirname, '../dist/js'),
    filename: '[name].js',
  },
  module: {
    rules: [
      {
        test: /\.(js)$/,
        exclude: /node_modules/,
        use: {
          loader: 'babel-loader',
          options: {
            presets: ['@babel/preset-env']
          }
        }
      },
      {
        test: /\.s(c|a)ss$/i,
        use: [
          //Create standalone css files
          //MiniCssExtractPlugin.loader,
          // Creates `style` nodes from JS strings
          // "style-loader",
          // Translates CSS into CommonJS
          'vue-style-loader',
          // Compiles Sass to CSS
          'css-loader',
          {
            loader: 'sass-loader',
            // Requires sass-loader@^7.0.0
            options: {
              implementation: require('sass'),
              indentedSyntax: true // optional
            },
            // Requires >= sass-loader@^8.0.0
            options: {
              implementation: require('sass'),
              sassOptions: {
                indentedSyntax: true // optional
              },
            },
          },
        ],
      },
      {
        test: /\.svg$/,
        use: 'svgo-loader',
        type: 'asset'
      },
      {
        test: /\.vue$/,
        use: 'vue-loader'
      },
      /*{ test: /\.css$/, use: ['vue-style-loader', 'css-loader'] },*/
      {
        test: /\.tsx?$/,
        use: 'ts-loader',
        exclude: /node_modules/,
      },
    ]
  },
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
      assets: path.resolve(__dirname, 'img')
    },
  },
  plugins: [
    new VueLoaderPlugin(),
    new webpack.ProvidePlugin({
      $: "jquery",
      jquery: "jQuery",
      "window.jQuery": "jquery"
    }),
  /*  new MiniCssExtractPlugin({
      filename: '../css/[name].css',
    }),*/
    new CompressionPlugin({
      include: /\.js$/,
      deleteOriginalAssets: true,
    }),
  ],
  externals: {
    jquery: 'jQuery',
    OC: "OC"
  },
};