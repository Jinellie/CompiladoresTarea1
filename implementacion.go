package main

import "fmt"

// Stack representa una pila (LIFO) ----------------------------------------------------------------
type Stack struct {
        items []interface{} // Slice para almacenar los elementos de la pila
}

// Push agrega un elemento a la cima de la pila
func (s *Stack) Push(item interface{}) {
        s.items = append(s.items, item)
}

// Pop elimina y devuelve el elemento de la cima de la pila
func (s *Stack) Pop() interface{} {
        if s.IsEmpty() {
                return nil // Pila vacía
        }
        lastIndex := len(s.items) - 1
        item := s.items[lastIndex]
        s.items = s.items[:lastIndex]
        return item
}

// Peek devuelve el elemento de la cima de la pila sin eliminarlo
func (s *Stack) Peek() interface{} {
        if s.IsEmpty() {
                return nil // Pila vacía
        }
        return s.items[len(s.items)-1]
}

// IsEmpty verifica si la pila está vacía
func (s *Stack) IsEmpty() bool {
        return len(s.items) == 0
}

// Size devuelve el número de elementos en la pila
func (s *Stack) Size() int {
        return len(s.items)
}

// Queue representa una fila (FIFO) ----------------------------------------------------------------
type Queue struct {
        items []interface{} // Slice para almacenar los elementos de la fila
}

// Enqueue agrega un elemento al final de la fila
func (q *Queue) Enqueue(item interface{}) {
        q.items = append(q.items, item)
}

// Dequeue elimina y devuelve el elemento del frente de la fila
func (q *Queue) Dequeue() interface{} {
        if q.IsEmpty() {
                return nil // Fila vacía
        }
        item := q.items[0]
        q.items = q.items[1:]
        return item
}

// PeekFront devuelve el elemento del frente de la fila sin eliminarlo
func (q *Queue) PeekFront() interface{} {
        if q.IsEmpty() {
                return nil // Fila vacía
        }
        return q.items[0]
}

// IsEmpty verifica si la fila está vacía
func (q *Queue) IsEmpty() bool {
        return len(q.items) == 0
}

// Size devuelve el número de elementos en la fila
func (q *Queue) Size() int {
        return len(q.items)
}

// Dictionary representa un diccionario ----------------------------------------------------------------
type Dictionary struct {
        items map[interface{}]interface{} // Mapa para almacenar pares clave-valor
}

// NewDictionary crea un nuevo diccionario
func NewDictionary() *Dictionary {
        return &Dictionary{items: make(map[interface{}]interface{})}
}

// Set agrega o actualiza un par clave-valor en el diccionario
func (d *Dictionary) Set(key, value interface{}) {
        d.items[key] = value
}

// Get devuelve el valor asociado a una clave
func (d *Dictionary) Get(key interface{}) interface{} {
        return d.items[key]
}

// Delete elimina un par clave-valor del diccionario
func (d *Dictionary) Delete(key interface{}) {
        delete(d.items, key)
}

// Contains verifica si una clave existe en el diccionario
func (d *Dictionary) Contains(key interface{}) bool {
        _, ok := d.items[key]
        return ok
}

// Size devuelve el número de elementos en el diccionario
func (d *Dictionary) Size() int {
        return len(d.items)
}

// Función main con pruebas de funcionamiento ----------------------------------------------------------------
func main() {
        // Pila (Stack)
        stack := Stack{}

        // Caso de prueba: Pila vacía
        fmt.Println("Caso de prueba: Pila vacía")
        fmt.Println("Salida esperada: true")
        fmt.Println("Salida obtenida:", stack.IsEmpty())
        fmt.Println("--------------------")

        // Caso de prueba: Agregar elementos y verificar el orden LIFO
        fmt.Println("Caso de prueba: Agregar elementos y verificar el orden LIFO")
        stack.Push(1)
        stack.Push(2)
        stack.Push(3)
        fmt.Println("Salida esperada: [1 2 3]")
        fmt.Println("Salida obtenida:", stack.items)
        fmt.Println("--------------------")

        // Caso de prueba: Eliminar elementos y verificar el orden LIFO
        fmt.Println("Caso de prueba: Eliminar elementos y verificar el orden LIFO")
        fmt.Println("Salida esperada: 3")
        fmt.Println("Salida obtenida:", stack.Pop())
        fmt.Println("Salida esperada: [1 2]")
        fmt.Println("Salida obtenida:", stack.items)
        fmt.Println("--------------------")

        // Caso de prueba: Peek (ver el elemento superior sin eliminarlo)
        fmt.Println("Caso de prueba: Peek (ver el elemento superior sin eliminarlo)")
        fmt.Println("Salida esperada: 2")
        fmt.Println("Salida obtenida:", stack.Peek())
        fmt.Println("--------------------")

        // Caso de prueba: Tamaño de la pila
        fmt.Println("Caso de prueba: Tamaño de la pila")
        fmt.Println("Salida esperada: 2")
        fmt.Println("Salida obtenida:", stack.Size())
        fmt.Println("--------------------")

        // Fila (Queue)
        queue := Queue{}

        // Caso de prueba: Fila vacía
        fmt.Println("Caso de prueba: Fila vacía")
        fmt.Println("Salida esperada: true")
        fmt.Println("Salida obtenida:", queue.IsEmpty())
        fmt.Println("--------------------")

        // Caso de prueba: Agregar elementos y verificar el orden FIFO
        fmt.Println("Caso de prueba: Agregar elementos y verificar el orden FIFO")
        queue.Enqueue("a")
        queue.Enqueue("b")
        queue.Enqueue("c")
        fmt.Println("Salida esperada: [a b c]")
        fmt.Println("Salida obtenida:", queue.items)
        fmt.Println("--------------------")

        // Caso de prueba: Eliminar elementos y verificar el orden FIFO
        fmt.Println("Caso de prueba: Eliminar elementos y verificar el orden FIFO")
        fmt.Println("Salida esperada: a")
        fmt.Println("Salida obtenida:", queue.Dequeue())
        fmt.Println("Salida esperada: [b c]")
        fmt.Println("Salida obtenida:", queue.items)
        fmt.Println("--------------------")

        // Caso de prueba: PeekFront (ver el elemento del frente sin eliminarlo)
        fmt.Println("Caso de prueba: PeekFront (ver el elemento del frente sin eliminarlo)")
        fmt.Println("Salida esperada: b")
        fmt.Println("Salida obtenida:", queue.PeekFront())
        fmt.Println("--------------------")

        // Caso de prueba: Tamaño de la fila
        fmt.Println("Caso de prueba: Tamaño de la fila")
        fmt.Println("Salida esperada: 2")
        fmt.Println("Salida obtenida:", queue.Size())
        fmt.Println("--------------------")

        // Diccionario (Dictionary)
        dict := NewDictionary()

        // Caso de prueba: Diccionario vacío
        fmt.Println("Caso de prueba: Diccionario vacío")
        fmt.Println("Salida esperada: true")
        fmt.Println("Salida obtenida:", dict.Size() == 0)
        fmt.Println("--------------------")

        // Caso de prueba: Agregar y obtener elementos
        fmt.Println("Caso de prueba: Agregar y obtener elementos")
        dict.Set("nombre", "Juan")
        dict.Set("edad", 30)
        fmt.Println("Salida esperada: map[nombre:Juan edad:30]")
        fmt.Println("Salida obtenida:", dict.items)
        fmt.Println("Salida esperada: Juan")
        fmt.Println("Salida obtenida:", dict.Get("nombre"))
        fmt.Println("--------------------")

        // Caso de prueba: Eliminar un elemento
        fmt.Println("Caso de prueba: Eliminar un elemento")
        dict.Delete("edad")
        fmt.Println("Salida esperada: map[nombre:Juan]")
        fmt.Println("Salida obtenida:", dict.items)
        fmt.Println("--------------------")

        //Caso de prueba: verificar si existe una llave.
        fmt.Println("Caso de prueba: verificar si existe una llave.")
        fmt.Println("Salida esperada: true")
        fmt.Println("Salida obtenida:", dict.Contains("nombre"))
        fmt.Println("Salida esperada: false")
        fmt.Println("Salida obtenida:", dict.Contains("edad"))
        fmt.Println("--------------------")

        // Caso de prueba: Tamaño del diccionario
        fmt.Println("Caso de prueba: Tamaño del diccionario")
        fmt.Println("Salida esperada: 1")
        fmt.Println("Salida obtenida:", dict.Size())
        fmt.Println("--------------------")
}