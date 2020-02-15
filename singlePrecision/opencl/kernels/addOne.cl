__kernel void addOne(__global float* data) {
	const int i = get_global_id (0);
	data[i] += 1;
}