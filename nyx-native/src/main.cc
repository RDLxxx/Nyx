#include <napi.h>

#include <filesystem>
#include <fstream>
#include <iostream>
#include <string>
#include <vector>

int AddFolder() {
    std::filesystem::path conf_path = "nyxconf";
    std::filesystem::path ses_path = "nyxconf/sessions";

    std::filesystem::create_directory(conf_path);
    std::filesystem::create_directory(ses_path);
    return 0;
}

Napi::Value AddFile(const Napi::CallbackInfo& info) {
    Napi::Env env = info.Env();

    if (std::filesystem::exists("nyxconf/filename.txt")) {
        return Napi::String::New(env, "exist");
    } else {
        AddFolder();

        std::ofstream MyFile("nyxconf/filename.txt");
        
        MyFile << "test";
        MyFile.close();
            
        return Napi::String::New(env, "created");
    }  
}

Napi::Value ReadFile(const Napi::CallbackInfo& info) {
    Napi::Env env = info.Env();

    std::string myText;
    std::string result;

    std::string path = info[0].As<Napi::String>();

    std::ifstream MyReadFile(path);
        
    while (getline(MyReadFile, myText)) {
        result += myText + "\n";
    }

    MyReadFile.close();

    return Napi::String::New(env, result);
}

Napi::Object Init(Napi::Env env, Napi::Object exports) {
    exports.Set("addFile", Napi::Function::New(env, AddFile));
    exports.Set("readFile", Napi::Function::New(env, ReadFile));
    
    return exports;
}

NODE_API_MODULE(nativeaddon, Init)