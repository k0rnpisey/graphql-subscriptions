CREATE MIGRATION m1tjacrsa6yiduqvxd5zb6a3pd44dxjggka7zoe6a6csfjt2uu5brq
    ONTO m14uegwrocuhlvfoee2q5pa6sc4gggq4fwobjprdbnwrejlsbj6h4a
{
  ALTER TYPE default::Post {
      ALTER LINK author {
          SET REQUIRED USING (<default::User>{});
      };
  };
};
