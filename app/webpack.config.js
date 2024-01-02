const path = require('path');
const MiniCssExtractPlugin = require('mini-css-extract-plugin');
const TerserPlugin = require('terser-webpack-plugin');
const CssMinimizerPlugin = require("css-minimizer-webpack-plugin");
const ImageMinimizerPlugin = require("image-minimizer-webpack-plugin");
const CopyPlugin = require('copy-webpack-plugin');

module.exports = {
  mode: 'development', // 'production' for minified code
  entry: './src/index.ts',
  module: {
    rules: [
      {
        test: /\.ts$/,
        use: 'ts-loader',
        exclude: /node_modules/
      },
      {
        test: /\.css$/,
        use: [
          MiniCssExtractPlugin.loader,
          'css-loader',
          {
            loader: 'postcss-loader', // Ensure PostCSS is properly configured
            options: {
              postcssOptions: {
                config: path.resolve(__dirname, 'postcss.config.js'), // Point to your PostCSS config
              },
            },
          },
        ]
      },
      {
        test: /\.(jpe?g|png|gif)$/i,
        type: 'asset/resource',
        use: [
          {
            loader: ImageMinimizerPlugin.loader,
            options: {
              minimizer: {
                implementation: ImageMinimizerPlugin.sharpGenerate,
                options: {
                  encodeOptions: {
                    webp: {
                      quality: 90,
                    },
                  },
                },
              },
            },
          },
        ],
        generator: {
          filename: '../img/generated/[name][ext]' // Output WebP images to 'img/generated' folder
        }
      },

      {
        test: /\.svg$/,
        type: 'asset/resource',
        generator: {
          filename: '../img/[name][ext][query]' // Output SVGs to 'img' folder
        }
      },
    ]
  },
  optimization: {
    minimizer: [
      new TerserPlugin(), // Minimize JavaScript
      new CssMinimizerPlugin(), // Minimize CSS
    ],
  },
  resolve: {
    extensions: ['.ts', '.js']
  },
  plugins: [
    new MiniCssExtractPlugin({
      filename: '../css/bundle.css' // Adjusted path
    }),
    new CopyPlugin({ 
      patterns: [
        { from: "src/img", to: "../img/", }
      ]
  })
  ],
  output: {
    filename: 'bundle.js',
    path: path.resolve(__dirname, '../app/assets/dist/js') // Adjusted path
  }
};
