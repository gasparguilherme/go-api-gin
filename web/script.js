const API_URL = 'http://localhost:8000';

document.addEventListener('DOMContentLoaded', loadProducts);

document.getElementById('productForm').addEventListener('submit', createProduct);

async function loadProducts() {
    try {
        const response = await fetch(`${API_URL}/products`);
        const products = await response.json();
        
        const tbody = document.getElementById('productsList');
        tbody.innerHTML = '';
        
        products.forEach(product => {
            const row = tbody.insertRow();
            row.innerHTML = `
                <td>${product.id}</td>
                <td>${product.name}</td>
                <td>R$ ${product.price.toFixed(2)}</td>
                <td>
                    <button class="delete-btn" onclick="deleteProduct(${product.id})">
                        Excluir
                    </button>
                </td>
            `;
        });
    } catch (error) {
        console.error('Erro ao carregar produtos:', error);
        alert('Erro ao carregar produtos!');
    }
}

async function createProduct(event) {
    event.preventDefault();
    
    const name = document.getElementById('name').value;
    const price = parseFloat(document.getElementById('price').value);
    
    try {
        const response = await fetch(`${API_URL}/product`, {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json'
            },
            body: JSON.stringify({ name, price })
        });
        
        if (response.ok) {
            document.getElementById('productForm').reset();
            loadProducts();
            alert('Produto criado com sucesso!');
        } else {
            alert('Erro ao criar produto!');
        }
    } catch (error) {
        console.error('Erro:', error);
        alert('Erro ao criar produto!');
    }
}

async function deleteProduct(id) {
    if (!confirm('Tem certeza que deseja excluir este produto?')) {
        return;
    }
    
    try {
        const response = await fetch(`${API_URL}/product/${id}`, {
            method: 'DELETE'
        });
        
        if (response.ok) {
            loadProducts();
            alert('Produto excluído com sucesso!');
        } else {
            alert('Erro ao excluir produto!');
        }
    } catch (error) {
        console.error('Erro:', error);
        alert('Erro ao excluir produto!');
    }
}