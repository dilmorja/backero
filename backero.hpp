#include <iostream>
#include <cstdio>
#include <vector>

namespace back {

  struct UTI {
  public:
    std::string name;
    std::string version;
    std::string hosts;

    UTI(std::string name, std::string version, std::string hosts);
    ~UTI();
    std::string toString();
  };

  UTI::UTI(std::string name, std::string version, std::string hosts) {
    this->name = name;
    this->version = version;
    this->hosts = hosts;
  }

  UTI::~UTI() {
    delete &name;
    delete &version;
    delete &hosts;
  }

  std::string UTI::toString() {
    const char* prestring = "%s_%s_%s";

    int tmo = snprintf(nullptr, 0, prestring,
      this->name,
      this->version,
      this->hosts);

    std::vector<char> buffer(tmo+1);

    snprintf(&buffer[0], buffer.size(), prestring,
      this->name,
      this->version,
      this->hosts);

    std::string ready_string;

    for (char b: buffer) {
      ready_string.push_back(b);
    }

    return ready_string;
  }

}