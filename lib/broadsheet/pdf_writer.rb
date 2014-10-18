require "prawn"

class PDFWriter
  def self.write(articles, filename)
    Prawn::Document.generate(filename) do
      font_families.update("Chalet" => {
         :normal => "#{Dir.home}/Library/Fonts/Chalet Book.ttf",
         :italic => "#{Dir.home}/Library/Fonts/Chalet Book Italic.ttf",
         :bold => "#{Dir.home}/Library/Fonts/Chalet Book Bold.ttf",
         :bold_italic => "#{Dir.home}/Library/Fonts/Chalet Book Bold Italic.ttf",
       })

      font_families.update("DTL Flesichmann" => {
         :normal => "#{Dir.home}/Library/Fonts/DTL Fleischmann D Regular.ttf",
         :italic => "#{Dir.home}/Library/Fonts/DTL Fleischmann D Regular.ttf",
         :bold => "#{Dir.home}/Library/Fonts/DTL Fleischmann D Bold.ttf",
         :bold_italic => "#{Dir.home}/Library/Fonts/DTL Fleischmann D Bold Italic.ttf",
       })

      font "DTL Flesichmann"
      default_leading 2

      articles.each do |article|
        text article.title, size: 26, style: :bold_italic
        move_down 4
        text article.author, size: 16
        move_down 20

        column_box([0, cursor], columns: 2, width: bounds.width, reflow_margins: true) do
         text article.content, size: 10
        end

        start_new_page
      end
    end
  end
end
