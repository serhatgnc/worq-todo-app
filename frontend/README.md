# WORQ Todo App - Frontend

React + TypeScript + Tailwind ile geliştirilmiş, TDD metodolojisi kullanılarak inşa edilmiş todo uygulaması.

## 🏗️ Teknolojiler

- **React 18** + **TypeScript** - UI framework
- **Tailwind CSS** - Styling
- **Vite** - Build tool
- **Jest** + **React Testing Library** - Testing

## 📋 Özellikler

- ✅ Todo ekleme
- ✅ Duplicate todo engelleme  
- ✅ Boş todo engelleme
- ✅ Responsive design
- ✅ %100 test coverage

## 🚀 Kurulum ve Çalıştırma

### Gereksinimler
- Node.js 18+
- npm 9+

### Kurulum
```bash
cd frontend
npm install
```

### Development
```bash
npm run dev
# http://localhost:3000
```

### Test
```bash
# Tüm testleri çalıştır
npm test

# Watch mode
npm run test:watch

# Coverage raporu
npm test -- --coverage
```

### Build
```bash
npm run build
npm run preview
```

## 🧪 TDD Süreci

Bu proje **Test Driven Development (TDD)** metodolojisi ile geliştirilmiştir.

### TDD Döngüsü: Red → Green → Refactor

#### 1️⃣ Component Rendering
```typescript
// RED: Test yaz
test('should render TodoApp component', () => {
  render(<TodoApp />);
  expect(screen.getByTestId('todo-app')).toBeInTheDocument();
});

// GREEN: Minimum kod
export const TodoApp = () => <div data-testid="todo-app" />;

// REFACTOR: İyileştir
```

#### 2️⃣ UI Elements
```typescript
// RED: Input ve button test
test('should render input and add button', () => {
  render(<TodoApp />);
  expect(screen.getByPlaceholderText('Add a new todo')).toBeInTheDocument();
  expect(screen.getByRole('button', { name: 'Add' })).toBeInTheDocument();
});

// GREEN: UI elementleri ekle
// REFACTOR: Tailwind styling
```

#### 3️⃣ Input State Management
```typescript
// RED: Input'a yazma test
test('should allow user to type in the input field', () => {
  render(<TodoApp />);
  const input = screen.getByPlaceholderText('Add a new todo');
  fireEvent.change(input, { target: { value: 'Buy milk' } });
  expect(input.value).toBe('Buy milk');
});

// GREEN: useState ile controlled component
// REFACTOR: TypeScript interface
```

#### 4️⃣ Todo Functionality
```typescript
// RED: Todo ekleme test
test('should add todo when add button is clicked', () => {
  render(<TodoApp />);
  const input = screen.getByPlaceholderText('Add a new todo');
  const button = screen.getByRole('button', { name: 'Add' });
  
  fireEvent.change(input, { target: { value: 'Buy milk' } });
  fireEvent.click(button);
  
  expect(screen.getByText('Buy milk')).toBeInTheDocument();
});

// GREEN: Todo state + add functionality
// REFACTOR: Unique validation + styling
```

### Test Stratejisi

**Unit Tests**: Component davranışları  
**Integration Tests**: User interactions  
**Edge Cases**: Empty inputs, duplicates  

## 🏛️ Mimari Kararlar

### State Management
- **Basit useState**: Redux vb state kütüphaneleri gereksiz, local state yeterli
- **Controlled Components**: Form best practices
- **Immutable Updates**: Spread operator kullanımı

### TypeScript
```typescript
interface Todo {
  id: string;    // Unique identifier
  text: string;  // Todo content
}
```

### Styling
- **Tailwind CSS**: Utility-first, responsive design
- **Component-based**: Her component kendi styling'i

### Testing
- **React Testing Library**: User-centric testing
- **Jest**: Unit testing framework  
- **FireEvent**: User interaction simulation

## 📊 Test Coverage
Statements : 100% ( 45/45 )
Branches : 100% ( 12/12 )
Functions : 100% ( 8/8 )
Lines : 100% ( 40/40 )

## 🚧 Gelecek Geliştirmeler

- [ ] Backend API integration

## 🔗 API Integration (Planlanan)

Backend API'si tamamlandığında:

```typescript
// API service
const addTodo = async (text: string) => {
  const response = await fetch('/api/todos', {
    method: 'POST',
    headers: { 'Content-Type': 'application/json' },
    body: JSON.stringify({ text })
  });
  return response.json();
};
```