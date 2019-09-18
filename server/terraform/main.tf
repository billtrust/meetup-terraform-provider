provider "btmascot" {
  root_url = "localhost:3000"
}

resource "btmascot_mascot" "new_mascot" {
  name        = "mascot 1"
  description = "a new description 2 bill"
}

resource "btmascot_mascot" "second_mascot" {
  name        = "mascot 2"
  description = "${btmascot_mascot.new_mascot.name} billtrust"
}
