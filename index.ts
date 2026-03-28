const myCar = {
  id: 1,
  manufacturer: "Toyota",
  model: "Camry",
  year: 2022,
  condition: "новая",
  price: 2500000
};

console.log("🚗 Наша машина:");
console.log("Производитель:", myCar.manufacturer);
console.log("Модель:", myCar.model);
console.log("Год:", myCar.year);
console.log("Состояние:", myCar.condition);
console.log("Цена:", myCar.price, "₽");

myCar.price = 2400000;
console.log("\n✅ Цена обновлена!");
console.log("Новая цена:", myCar.price, "₽");

console.log("\n📌 Обновлённая карточка:");
console.log(myCar);