using UnityEngine;
using System.Collections;
using TMPro;

public class HealthBar : MonoBehaviour
{
    public TextMeshPro health;
    private int healthValue;
    // Use this for initialization

    void Start()
    {

    }

    // Update is called once per frame
    void Update()
    {

    }
    public void reduceByOne()
    {
        setHealth(healthValue -1);
    }
    // Used to set the health
    public void setHealth(int hearts)
    {
        health.text = hearts.ToString();
        healthValue = hearts;
    }
}
