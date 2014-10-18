require 'broadsheet/sources/disegno-daily'

class Builder
  def build
    DisegnoDaily.download
  end
end