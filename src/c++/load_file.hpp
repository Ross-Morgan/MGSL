#pragma once

#include <sys/stat.h>
#include <cstdlib>
#include <fstream>
#include <iostream>
#include <string>


bool
fileExists(const std::string& fileName)
{
    struct stat buffer;
    return (stat (fileName.c_str(), &buffer) == 0);
}


std::string
loadFile(const std::string& fileName)
{
    if (!fileExists(fileName))
    {
        std::cerr << "File " << fileName << " does not exist.";
        exit(1);
    }

    std::string fileContents{};
    std::string line;
    std::ifstream file(fileName.c_str());

    while (getline(file, line))
    {
        fileContents.append(line);
        fileContents.append("\n");
    }

    std::cout << fileContents << std::endl;

    return fileContents;
}
