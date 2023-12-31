const path = require('path');
const MiniCssExtractPlugin = require('mini-css-extract-plugin');

module.exports = {
  mode: 'development', // 'production' for minified code
  entry: './src/index.ts',
  optimization: { usedExports: false },
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
      }
    ]
  },
  resolve: {
    extensions: ['.ts', '.js']
  },
  plugins: [
    new MiniCssExtractPlugin({
      filename: '../css/bundle.css' // Adjusted path
    })
  ],
  output: {
    filename: 'bundle.js',
    path: path.resolve(__dirname, '../app/assets/dist/js') // Adjusted path
  }
};
