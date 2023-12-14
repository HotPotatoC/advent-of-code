#include <bits/stdc++.h>

auto parse_input() {
  std::ifstream input;

  // input.open("input.in");
  input.open("input.sample.in");

  std::vector<std::vector<int>> grid;

  std::string line;

  while (std::getline(input, line)) {
    std::vector<int> row;

    for (auto c : line) {
      row.push_back(c - '0');
    }

    grid.push_back(row);
  }

  return 0;
}

auto part_one() {
  auto input = parse_input();
  //
  return 0;
}

auto part_two() {
  auto input = parse_input();
  //
  return 0;
}

int main() {
  std::cout << "Part one: " << part_one() << std::endl;
  std::cout << "Part two: " << part_two() << std::endl;

  return 0;
}