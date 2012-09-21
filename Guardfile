# -*- coding: utf-8 -*-
# Coffeescriptの監視
guard 'coffeescript', :input => 'coffee', :output => 'dist/js', :bare => true

# SCSSの監視
guard :shell do
  # import以下が編集されたら、main.scssを更新したことにする
  watch(/^(.+\/)*.+(\.scss)$/) {
    |m|
        `mkdir -p dist/css`
        `sass --scss scss/main.scss dist/css/main.css`
  }
end

# HTMLテンプレートファイルの監視
guard :shell do
  watch(/^templates\/html\/(.+\/)*.+(\.html|\.json)$/) {
    |m|
    `mkdir -p dist`
    `./template -p "templates/html/*.html" -t "index" -f templates/html/data.json > dist/index.html`
 }
end

# CoffeeScriptテンプレートファイルの監視
guard :shell do
  watch(/^templates\/coffee\/(((.+\/)*.+)(\.coffee|\.json)$)/) {
    |m|
    dist = "coffee/" + m[1]
    dir = `dirname #{dist}`
    `mkdir -p #{dir}`
    `./template -p "#{m[0]}" -t "#{m[2]}" -f templates/coffee/data.json > #{dist}`
 }
end


# ライブラリの監視
guard :shell do
  watch(/^lib\/(.+\/)*.+(\.js|\.css|\.jpg|\.png)$/) {
    |m|
    dist = "dist/" + m[0]
    dir = `dirname #{dist}`
    `mkdir -p #{dir}`
    `cp #{m[0]} #{dist}`
  }
end
