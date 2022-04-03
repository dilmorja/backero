#include <iostream>
#include <cstdio>

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
    std::string ready_string;
    char* buf;

    sprintf(buf, "%s_%s_%s", this->name, this->version, this->hosts);

    ready_string.assign(buf, sizeof(buf));

    return ready_string;
  }

}