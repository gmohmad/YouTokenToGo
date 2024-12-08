#ifdef __cplusplus
extern "C" {
#endif
	void* new_base_encoder(const char* model_path, int n_threads);
	void destroy_base_encoder(void* encoder_ptr);
	int encode_as_ids(void* encoder_ptr, const char* sentence, int** output_ids, int* length);
	void free_ids(int* ids);

#ifdef __cplusplus
}
#endif
