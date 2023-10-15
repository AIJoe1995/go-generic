package mapx

// key value 对key求hash值 把key value从在对应hash值的map处， 因为可能会出现hash冲突，使用一个单链表解决hash冲突

type node[K Hashable, V any] struct {
	key   K
	value V
	next  *node[K, V]
}

func newNode[K Hashable, V any](key K, value V) *node[K, V] {
	return &node[K, V]{
		key:   key,
		value: value,
	}
}

type Hashable interface {
	// Code 返回该元素的哈希值
	// 注意：哈希值应该尽可能的均匀以避免冲突
	Code() uint64
	// Equals 比较两个元素是否相等。如果返回 true，那么我们会认为两个键是一样的。
	Equals(key any) bool
}

type HashMap[K Hashable, V any] struct {
	hashmap map[uint64]*node[K, V]
}

func NewHashMap[K Hashable, V any](size int) *HashMap[K, V] {
	return &HashMap[K, V]{
		hashmap: make(map[uint64]*node[K, V], size),
	}
}

func (h *HashMap[K, V]) Put(key K, value V) error {
	hash := key.Code()
	node, ok := h.hashmap[hash]
	if !ok {
		newnode := newNode[K, V](key, value)
		h.hashmap[hash] = newnode
		return nil
	}
	root := node
	prev := node
	for root != nil {
		if root.key.Equals(key) {
			root.value = value
			return nil
		}
		prev = root
		root = root.next
	}
	newnode := newNode[K, V](key, value)
	prev.next = newnode
	return nil

}

func (h *HashMap[K, V]) Get(key K) (V, bool) {
	var zero V
	var hash = key.Code()
	node, ok := h.hashmap[hash]
	if !ok {
		return zero, false
	}
	root := node
	for root != nil {
		if root.key.Equals(key) {
			return root.value, true
		}
		root = root.next
	}
	return zero, false
}
