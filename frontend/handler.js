exports.render = async function(event) {
      const response = {
        statusCode: 200,
        body: JSON.stringify({
            foo: 'Hello, world',
            bar: 'Goodbye, world',
        })
    };

    return response;
}
