import { useParams } from "react-router-dom";
import { useQuery } from "@apollo/client";
import { GET_MENU_ITEM_BY_ID} from "../graphql/Menuqueries";
import Navbar from "../components/NavBar";






const MenuItem: React.FC = () => {
    const { menuId } = useParams();

    // ! this is to fetch the menu ITEM BY ID
    const { data, loading, error } = useQuery(GET_MENU_ITEM_BY_ID, {
        variables: { id: menuId },
        skip: !menuId,
    });

    
    


    // ! initialize mutation for delete a menu item
    const [deleteMenuItem] = useMutation(DELETE_MENU_ITEM); // 🔹 Added delete mutation


    if (loading) return <p>Loading...</p>;
    if (error) return <p>Error fetching menu data: {error.message}</p>;
    if (!data || !data.getMenuItemById) return <p>No menu data found.</p>;

    const item = data.getMenuItemById;

    
    return (
        <>
    <Navbar />
    <div className="flex justify-center items-center min-h-[90vh] bg-gray-100"> {/* 🔹 Center container */}
        <div className="p-8 w-full max-w-2xl bg-white rounded-xl shadow-md space-y-6 max-h-[85vh] overflow-y-auto"> {/* 🔹 Bigger card with scroll */}
            <button
            onClick={() => navigate("/dashboard")}
                onClick={() => navigate(-1)}
                className="text-blue-600 underline"
            >
                Back
            </button>

            <img src={item.image_url} alt={item.name} className="w-full h-72 object-cover rounded" /> {/* 🔹 Image with fixed height */}
            <h2 className="text-gray-900 text-3xl font-bold">{item.name}</h2> {/* 🔹 Slightly bigger title */}
            <p className="text-gray-700 text-lg">{item.description}</p>
            <p className="text-2xl font-semibold text-green-700">${item.price.toFixed(2)}</p>
            <p className="text-base text-gray-500">Category: {item.category}</p>
            <p className="text-base text-gray-500">
                Status: {item.availability_status ? "Available" : "Out of stock"}
            </p>

            <div className="flex justify-between">
                <button
                    onClick={() => navigate(`/menu/update/${menuId}`)}
                    className="bg-black text-yellow-500 px-4 py-2 rounded hover:bg-gray-800"
                >
                    Update
                </button>

                <button
                    onClick={async () => {
                        try {
                            await deleteMenuItem({ variables: { id: menuId } });
                            navigate("/menu");
                        } catch (err) {
                            alert("Failed to delete menu item.");
                        }
                    }}
                    className="bg-black text-red-500 px-4 py-2 rounded hover:bg-gray-800"
                >
                    Delete
                </button>
            </div>
        </div>
    </div>
</>


    );
};

export default MenuItem;