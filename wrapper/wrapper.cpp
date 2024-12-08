#include <vector>
#include <string>
#include "../src/bpe.h"
#include "../src/utils.h"

using namespace vkcom;

extern "C" {
	void* new_base_encoder(const char* model_path, int n_threads) {
		Status ret_status;
		BaseEncoder* encoder = new BaseEncoder(std::string(model_path), n_threads, &ret_status);
		if (!ret_status.ok()) {
			delete encoder;
			return nullptr;
		}
		return encoder;
	}

	void destroy_base_encoder(void* encoder_ptr) {
		delete static_cast<BaseEncoder*>(encoder_ptr);
	}

	int encode_as_ids(void* encoder_ptr, const char* sentence, int** output_ids, int* length) {
		auto* encoder = static_cast<BaseEncoder*>(encoder_ptr);
		std::vector<std::string> sentences = {std::string(sentence)};
		std::vector<std::vector<int>> ids;

		Status status = encoder->encode_as_ids(sentences, &ids);
		if (!status.ok() || ids.empty()) {
			return -1;
		}

		*length  = ids[0].size();
		*output_ids = new int[*length];
		for (int i = 0; i < *length; ++i) {
			(*output_ids)[i] = ids[0][i];
		}

		return 0;
	}
	
	void free_ids(int* ids) {
		delete[] ids;
	}
}
