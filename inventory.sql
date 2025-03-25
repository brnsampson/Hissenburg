CREATE TABLE inventory (
    inventory_id INTEGER PRIMARY KEY,
    left_hand INTEGER REFERENCES item(item_id),
    right_hand INTEGER REFERENCES item(item_id),
    head INTEGER REFERENCES item(item_id),
    torso INTEGER REFERENCES item(item_id),
    backpack_0 INTEGER REFERENCES item(item_id),
    backpack_1 INTEGER REFERENCES item(item_id),
    backpack_2 INTEGER REFERENCES item(item_id),
    backpack_3 INTEGER REFERENCES item(item_id),
    backpack_4 INTEGER REFERENCES item(item_id),
    backpack_5 INTEGER REFERENCES item(item_id),
    extra_space_0 INTEGER REFERENCES item(item_id),
    extra_space_1 INTEGER REFERENCES item(item_id),
    extra_space_2 INTEGER REFERENCES item(item_id),
    extra_space_3 INTEGER REFERENCES item(item_id),
    extra_space_4 INTEGER REFERENCES item(item_id),
    extra_space_5 INTEGER REFERENCES item(item_id),
    ground_0 INTEGER REFERENCES item(item_id),
    ground_1 INTEGER REFERENCES item(item_id),
    ground_2 INTEGER REFERENCES item(item_id),
    ground_3 INTEGER REFERENCES item(item_id),
    ground_4 INTEGER REFERENCES item(item_id),
    ground_5 INTEGER REFERENCES item(item_id)
);

