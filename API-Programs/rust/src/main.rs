use warp::Filter;
use rand::Rng; // Rng trait must be in scope to use the rand::thread_rng function

#[tokio::main]
async fn main() {
    // Define the route
    let route = warp::path::end().map(|| {
        let random_number = rand::thread_rng().gen_range(0..100).to_string();
        warp::reply::html(random_number)
    });

    println!("Starting server on :4040");
    warp::serve(route).run(([0, 0, 0, 0], 4040)).await;
}