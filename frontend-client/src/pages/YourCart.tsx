import { useQuery, useMutation } from "@apollo/client";
import { GET_CART_AND_ITEMS, CREATE_ORDER_FROM_CART, UPDATE_CART_ITEM, REMOVE_CART_ITEM } from "../graphql/Orderqueries";
import Navbar from "../components/NavBar";
import "../styles/main.css";
import { useNavigate, useParams } from "react-router-dom";

import React, { useEffect, useState } from "react";

const YourCart: React.FC = () => {
    const navigate = useNavigate();
    const { userId } = useParams<{ userId: string }>();

    const [orderType, setOrderType] = useState("dine-in");
    const [orderData, setOrderData] = useState<any>(null);
    const [updateErrors, setUpdateErrors] = useState<{ [key: string]: string }>({});
    const [deliveryAddress, setDeliveryAddress] = useState("");
    const [specialRequests, setSpecialRequests] = useState("");

    const { data, loading, error, refetch } = useQuery(GET_CART_AND_ITEMS, {
        variables: { user_id: userId },
        skip: !userId,
    });

    const [cartItems, setCartItems] = useState<any[]>([]);
    const cart = data?.getCartAndMenuItems;

    useEffect(() => {
        if (orderType === "dine-in" || orderType === "takeout") setDeliveryAddress("");
    }, [orderType]);

    useEffect(() => {
        if (cart?.items) {
            setCartItems(cart.items);
        }
    }, [cart]);

    const [createOrderFromCart, { loading: creatingOrder, error: createOrderError }] = useMutation(CREATE_ORDER_FROM_CART);
    const [updateCartItem] = useMutation(UPDATE_CART_ITEM);
    const [removeCartItem] = useMutation(REMOVE_CART_ITEM);

    const handleUpdate = async (item: any) => {
        try {
            await updateCartItem({
                variables: {
                    input: {
                        id: item.id,
                        menu_item_id: item.menu_item_id,
                        quantity: item.quantity,
                        customizations: item.customizations || "",
                    },
                },
            });
            setUpdateErrors(prev => ({ ...prev, [item.id]: "" }));
            await refetch();
        } catch (error) {
            console.error("Error updating cart item:", error);
            setUpdateErrors(prev => ({
                ...prev,
                [item.id]: (error instanceof Error ? error.message : "Failed to update item"),
            }));
        }
    };

    const handleCreateOrder = async () => {
        try {
            if (creatingOrder) return;
            const { data } = await createOrderFromCart({
                variables: {
                    cartID: cart.id,
                    userID: userId,
                    orderType,
                    deliveryAddress: orderType === "delivery" ? deliveryAddress : null,
                    specialRequests,
                },
            });
            console.log("Order created:", data.createOrderFromCart);
            setOrderData(data);
            navigate(`/payment/${data.createOrderFromCart.id}`);
        } catch (err) {
            console.error("Failed to create order:", err);
        }
    };

    const handleDelete = async (item: any) => {
        try {
            const { data } = await removeCartItem({ variables: { id: item.id } });
            if (data.removeCartItem) {
                setCartItems((prev) => prev.filter((i) => i.id !== item.id));
            } else {
                console.error("Item not found or failed to delete.");
            }
        } catch (error) {
            console.error("Error deleting cart item:", error);
        }
    };

    const totalPrice = cartItems.reduce((total: number, item: any) => {
        return total + item.price * item.quantity;
    }, 0).toFixed(2);

    if (loading) return <p>Loading your cart...</p>;
    if (error) return <p>Error loading cart: {error.message}</p>;

    if (!cartItems || cartItems.length === 0) {
        return (
            <>
                <Navbar />
                <div className="p-8 text-center">
                    <h2 className="text-2xl font-bold">Your Cart</h2>
                    <p className="text-gray-500 mt-4">Your cart is empty. Start ordering now!</p>
                    <p className="text-gray-500 mt-4">If not showing, try reloading!</p>
                </div>
            </>
        );
    }

    return (
        <>
            <Navbar />
            <div className="p-8 max-w-4xl mx-auto">
                <h2 className="text-3xl font-bold mb-6 text-black">Your Cart</h2>
                <div className="max-h-[500px] overflow-y-auto border rounded-lg p-4 shadow-xl bg-white scrollbar-hide">
                    {cartItems.map((item: any, index: number) => (
                        <div
                            key={item.id}
                            className="border-b last:border-none pb-3 mb-3 flex justify-between items-start gap-4"
                        >
                            <img
                                src={item.menuItem?.image_url}
                                alt={item.menuItem?.name}
                                className="w-30 h-30 object-cover rounded-md border"
                            />

                            <div className="flex-1 space-y-2">
                                <p className="text-lg font-semibold text-black">{item.menuItem?.name}</p>

                                <label className="block text-sm text-gray-600">
                                    Quantity:
                                    <input
                                        type="number"
                                        min={1}
                                        className="ml-2 border rounded px-2 py-1 w-20"
                                        value={item.quantity}
                                        onChange={(e) => {
                                            const newQty = parseInt(e.target.value, 10);
                                            setCartItems(prev =>
                                                prev.map((item, i) =>
                                                    i === index ? { ...item, quantity: newQty } : item
                                                )
                                            );
                                        }}
                                    />
                                </label>

                                <label className="block text-sm text-gray-600">
                                    Customizations:
                                    <input
                                        type="text"
                                        className="ml-2 border rounded px-2 py-1 w-full mt-1"
                                        value={item.customizations || ""}
                                        onChange={(e) => {
                                            setCartItems(prev =>
                                                prev.map((item, i) =>
                                                    i === index ? { ...item, customizations: e.target.value } : item
                                                )
                                            );
                                        }}
                                    />
                                </label>

                                <p className="text-sm text-gray-600">
                                    Added on: {new Date(item.created_at).toLocaleString()}
                                </p>
                                {updateErrors[item.id] && (() => {
                                    const match = updateErrors[item.id].match(/available: (\d+), requested: (\d+)/);
                                    if (updateErrors[item.id].toLowerCase().includes("not enough stock") && match) {
                                        const [_, available, requested] = match;
                                        return (
                                            <p className="text-red-500 text-sm mt-2">
                                                Only {available} items available. You requested {requested}.
                                            </p>
                                        );
                                    }
                                    return (
                                        <p className="text-red-500 text-sm mt-2">{updateErrors[item.id]}</p>
                                    );
                                })()}
                            </div>

                            <div className="text-right min-w-[80px]">
                                {item.price !== item.menuItem.price ? (
                                    <>
                                        <p className="text-lg text-gray-500 line-through">
                                            ₱{item.menuItem.price.toFixed(2)}
                                        </p>
                                        <p className="text-lg font-bold text-green-700">
                                            ₱{item.price.toFixed(2)}
                                        </p>
                                    </>
                                ) : (
                                    <p className="text-lg font-bold text-green-700">
                                        ₱{item.price.toFixed(2)}
                                    </p>
                                )}
                            </div>

                            <div className="flex flex-col gap-2">
                                <button
                                    onClick={() => handleUpdate(item)}
                                    className="bg-blue-600 text-white px-4 py-2 rounded hover:bg-blue-700"
                                >
                                    Update
                                </button>
                                <button
                                    onClick={() => handleDelete(item)}
                                    className="bg-red-600 text-white px-4 py-2 rounded hover:bg-red-700"
                                >
                                    Delete
                                </button>
                            </div>
                        </div>
                    ))}
                </div>
                <div className="mt-6 border-t pt-6">
                    <h3 className="text-xl text-black font-semibold mb-4">Total: ₱{totalPrice}</h3>
                </div>

                <div className="mt-6 border-t pt-6">
                    <form
                        className="fixed bottom-0 left-0 w-full bg-white shadow-lg p-4 flex flex-wrap gap-4 items-end border-t z-50"
                        onSubmit={(e) => {
                            e.preventDefault();
                            createOrderFromCart({
                                variables: {
                                    cartID: cart.id,
                                    userID: userId,
                                    orderType,
                                    deliveryAddress,
                                    specialRequests,
                                },
                            });
                        }}
                    >
                        <div className="flex-1 min-w-[200px]">
                            <label className="block text-sm font-semibold text-gray-600">Order Type</label>
                            <select
                                title="Order Type"
                                value={orderType}
                                onChange={(e) => setOrderType(e.target.value)}
                                className="w-full p-2 border rounded-lg text-black"
                            >
                                <option value="dine-in">Dine-in</option>
                                <option value="takeout">Takeout</option>
                                <option value="delivery">Delivery</option>
                            </select>
                        </div>

                        <div className="flex-1 min-w-[200px]">
                            <label className="block text-sm font-semibold text-gray-600">Delivery Address</label>
                            <input
                                type="text"
                                value={deliveryAddress}
                                onChange={(e) => setDeliveryAddress(e.target.value)}
                                className="w-full p-2 border rounded-lg text-black"
                                placeholder="e.g. 123 Cyber Lane"
                                disabled={orderType === "dine-in" || orderType === "takeout"}
                            />
                        </div>

                        <div className="flex-1 min-w-[200px]">
                            <label className="block text-sm font-semibold text-gray-600">Special Requests</label>
                            <input
                                type="text"
                                value={specialRequests}
                                onChange={(e) => setSpecialRequests(e.target.value)}
                                className="w-full p-2 border rounded-lg text-black"
                                placeholder="e.g., Allergies, etc."
                            />
                        </div>

                        <div className="flex items-end gap-2">
                            <button
                                type="submit"
                                onClick={handleCreateOrder}
                                className="bg-green-600 text-white px-6 py-2 rounded-lg hover:bg-green-700 transition"
                                disabled={creatingOrder}
                            >
                                {creatingOrder ? "Placing..." : "Place Order"}
                            </button>
                            <button
                                type="button"
                                onClick={() => window.location.reload()}
                                className="bg-gray-400 text-white px-6 py-2 rounded-lg hover:bg-gray-500 transition"
                            >
                                Cancel
                            </button>
                        </div>
                    </form>

                    {createOrderError && (
                        <p className="text-red-600 mt-4">Error: {createOrderError.message}</p>
                    )}
                    {orderData && (
                        <p className="text-green-700 font-medium mt-4">
                            Order placed! Order ID: {orderData.createOrderFromCart.id}
                        </p>
                    )}
                </div>
            </div>
        </>
    );
};

export default YourCart;