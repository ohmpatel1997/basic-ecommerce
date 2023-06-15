import React, { useState } from "react";
import axios from 'axios';

export default function App() {
    const [formData, setFormData] = useState({ name: "", category: "", sku: "" });
    const [searchData, setSearch] = useState({ search: "" });
    const [results, setResults] = useState();

    const handleChange = (event) => {
        const { name, value } = event.target;
        setFormData((prevFormData) => ({ ...prevFormData, [name]: value }));
    };

    const handleSearchChange = (event) => {
        setResults("")
        const { name, value } = event.target;
        setSearch((prevFormData) => ({ ...prevFormData, [name]: value }));
    };

    const handleSubmit = (e) => {
        e.preventDefault();
        setFormData((prevFormData) => ({ ...prevFormData, name: "", category: "", sku: "" }));

        axios.post("http://localhost:8080/api/v1/products", {
            name: formData.name,
            category: formData.category,
            sku: formData.sku
        }).then((response) => {
            if (response.status > 199 && response.status <= 299) {
                window.alert("Product created successfully");
            } else {
                window.alert("Product creation failed" + response.data);
            }
        }).catch((error) => {
            window.alert("error calling API: " + error)
        })
    };

    const handleSearch = async (e) => {
        e.preventDefault();
        try {
            setResults("");
            const resp = await axios.get("http://localhost:8080/api/v1/products/search?query=" + searchData.search)
            if (resp.status > 199 && resp.status <= 299) {
                console.log(JSON.stringify(resp.data,null, 2));
                setResults(JSON.stringify(resp.data,null, 2));
            } else {
                window.alert("search failed" + resp.data);
            }
            setSearch((prevFormData) => ({ ...prevFormData, search: "" }));

        }catch(error) {
            window.alert("error calling API: " + error);
        }
    };

    return (
        <div>
            <br />
            <br />
            <br />
            CREATE PRODUCT
            <br />
            <br />
            <form onSubmit={handleSubmit}>
                <label htmlFor="name">Name:</label>
                <input
                    type="text"
                    id="name"
                    name="name"
                    value={formData.name}
                    onChange={handleChange}
                />

                <label htmlFor="category">Category:</label>
                <input
                    type="text"
                    id="category"
                    name="category"
                    value={formData.category}
                    onChange={handleChange}
                />

                <label htmlFor="sku">SKU:</label>
                <input
                    type="text"
                    id="sku"
                    name="sku"
                    value={formData.sku}
                    onChange={handleChange}
                />
                <button type="submit">Submit</button>

                <br />
                <br />
                <br />
            </form>
            SEARCH PRODUCT
            <br />
            <br />
            <input
                type="text"
                id="search"
                name="search"
                value={searchData.search}
                onChange={handleSearchChange}
            />
            <button onClick={handleSearch}>Search</button>
            <br />
            <br />
            <br />

            RESULTS
            <br />
            <br />
            <textarea
                id="result"
                rows="10"
                cols="100"
                value={results || ""}
            ></textarea>
        </div>
    );
}