require "prawn"

class PDFWriter

  # Given an array of Article objects and a filename, outputs a PDF broadsheet
  # for the articles to the specified filename
  def self.write(articles, filename)
    Prawn::Document.generate(filename) do
      font_families.update("Chalet" => {
         normal: "/Users/gklitt/Library/Fonts/Chalet Book.ttf",
         italic: "/Users/gklitt/Library/Fonts/Chalet Book Italic.ttf",
         bold: "/Users/gklitt/Library/Fonts/Chalet Book Bold.ttf",
         bold_italic: "/Users/gklitt/Library/Fonts/Chalet Book Bold Italic.ttf",
       })

      font_families.update("DTL Fleischmann" => {
         normal: "/Users/gklitt/Library/Fonts/DTL Fleischmann D Regular.ttf",
         italic: "/Users/gklitt/Library/Fonts/DTL Fleischmann D Medium Italic.ttf",
         bold: "/Users/gklitt/Library/Fonts/DTL Fleischmann D Bold.ttf",
         bold_italic: "/Users/gklitt/Library/Fonts/DTL Fleischmann D Bold Italic.ttf",
       })

      font "Chalet"
      text "Broadsheet", size: 60, style: :bold_italic
      move_down 6
      text Time.now.strftime("%B %-d, %Y"), size: 20, style: :italic
      move_down 4
      horizontal_rule

      font "DTL Fleischmann"
      default_leading 2

      bounding_box([250, 300], width: 280, height: 300) do
        text "Table of contents", size: 20, style: :italic
        move_down 10
        articles.each do |article|
          font "Chalet"
          text article.source, size: 10

          font "DTL Fleischmann"
          move_down 3
          text article.title, size: 14
          move_down 14
        end
      end

      articles.each do |article|
        start_new_page

        text article.title, size: 26, style: :bold_italic
        move_down 4
        text article.author, size: 16
        move_down 20

        column_box([0, cursor], columns: 2, width: bounds.width, reflow_margins: true) do
         text article.content, size: 10
        end
      end
    end
  end
end
