UPDATE products SET (
    	title=$1,
		description=$2,
		price=$3,
		ImgUrl=$4
) where id=$5;