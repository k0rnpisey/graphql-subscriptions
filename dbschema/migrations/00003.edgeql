CREATE MIGRATION m14uegwrocuhlvfoee2q5pa6sc4gggq4fwobjprdbnwrejlsbj6h4a
    ONTO m1w24s3ath5czb5pnjzp2h47xd66vf36hw3ygodlljxitxrugg4d5a
{
  ALTER TYPE default::Notification {
      ALTER LINK user {
          SET REQUIRED USING (<default::User>{});
      };
  };
};
